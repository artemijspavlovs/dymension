package delayedack

import (
	"errors"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"

	commontypes "github.com/dymensionxyz/dymension/v3/x/common/types"
	"github.com/dymensionxyz/dymension/v3/x/delayedack/keeper"
	"github.com/dymensionxyz/dymension/v3/x/delayedack/types"
	rollapptypes "github.com/dymensionxyz/dymension/v3/x/rollapp/types"
)

// IBCMiddleware implements the ICS26 callbacks
type IBCMiddleware struct {
	porttypes.IBCModule
	keeper.Keeper
	rollappKeeper types.RollappKeeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application
func NewIBCMiddleware(app porttypes.IBCModule, k keeper.Keeper, rollappKeeper types.RollappKeeper) IBCMiddleware {
	return IBCMiddleware{
		IBCModule:     app,
		Keeper:        k,
		rollappKeeper: rollappKeeper,
	}
}

// OnRecvPacket handles the receipt of a packet and puts it into a pending queue
// until its state is finalized
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	if !im.IsRollappsEnabled(ctx) {
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}
	logger := ctx.Logger().With(
		"module", types.ModuleName,
		"packet_source", packet.SourcePort,
		"packet_destination", packet.DestinationPort,
		"packet_sequence", packet.Sequence)

	rollappPortOnHub, rollappChannelOnHub := packet.DestinationPort, packet.DestinationChannel

	rollapp, transferPacketData, err := im.rollappKeeper.ExtractRollappAndTransferPacketFromData(
		ctx,
		packet.Data,
		rollappPortOnHub,
		rollappChannelOnHub,
	)
	if err != nil {
		logger.Error("Failed to extract rollapp id from packet", "err", err)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	rollappID := rollapp.RollappId

	if rollappID == "" {
		logger.Debug("Skipping IBC transfer OnRecvPacket for non-rollapp chain")
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	err = im.ValidateRollappId(ctx, rollappID, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		logger.Error("Failed to validate rollappID", "rollappID", rollappID, "err", err)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	proofHeight, err := im.GetProofHeight(ctx, commontypes.RollappPacket_ON_RECV, rollappPortOnHub, rollappChannelOnHub, packet.Sequence)
	if err != nil {
		logger.Error("Failed to get proof height from packet", "err", err)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	finalized, err := im.CheckIfFinalized(ctx, rollappID, proofHeight)
	if err != nil {
		logger.Error("Failed to check if packet is finalized", "err", err)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	if finalized {
		logger.Debug("Skipping eIBC transfer OnRecvPacket as the packet proof height is already finalized")
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	// Save the packet data to the store for later processing
	rollappPacket := commontypes.RollappPacket{
		RollappId:   rollappID,
		Packet:      &packet,
		Status:      commontypes.Status_PENDING,
		Relayer:     relayer,
		ProofHeight: proofHeight,
		Type:        commontypes.RollappPacket_ON_RECV,
	}
	im.SetRollappPacket(ctx, rollappPacket)

	logger.Debug("Set rollapp packet",
		"rollappID", rollappPacket.RollappId,
		"proofHeight", rollappPacket.ProofHeight,
		"type", rollappPacket.Type)

	err = im.eIBCDemandOrderHandler(ctx, rollappPacket, *transferPacketData)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	return nil
}

func unwrapPacketData(
	wrappedPacket channeltypes.Packet,
	transferPacketData *transfertypes.FungibleTokenPacketData,
) (packet channeltypes.Packet, err error) {
	wrappedPacket.Data, err = types.ModuleCdc.MarshalJSON(transferPacketData)
	packet = wrappedPacket
	return
}

// OnAcknowledgementPacket implements the IBCMiddleware interface
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	wrappedPacket channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	logger := ctx.Logger().With(
		"module", types.ModuleName,
		"packet_source", wrappedPacket.SourcePort,
		"packet_destination", wrappedPacket.DestinationPort,
		"packet_sequence", wrappedPacket.Sequence)

	rollappPortOnHub, rollappChannelOnHub := wrappedPacket.SourcePort, wrappedPacket.SourceChannel

	rollapp, transferPacketData, err := im.rollappKeeper.ExtractRollappAndTransferPacketFromData(ctx, wrappedPacket.Data, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		logger.Error("Failed to extract rollapp id from channel", "err", err)
		return err
	}

	rollappID := rollapp.RollappId

	packet, err := unwrapPacketData(wrappedPacket, transferPacketData)
	if err != nil {
		logger.Error("Failed to marshal transfer packet data", "err", err)
		return err
	}

	if rollappID == "" {
		logger.Debug("Skipping IBC transfer OnAcknowledgementPacket for non-rollapp chain")
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}

	err = im.ValidateRollappId(ctx, rollappID, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		logger.Error("Failed to validate rollappID", "rollappID", rollappID, "err", err)
		return err
	}

	if !im.IsRollappsEnabled(ctx) {
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}

	var ack channeltypes.Acknowledgement
	if err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		logger.Error("Unmarshal acknowledgement", "err", err)
		return errorsmod.Wrapf(types.ErrUnknownRequest, "unmarshal ICS-20 transfer packet acknowledgement: %v", err)
	}

	proofHeight, err := im.GetProofHeight(ctx, commontypes.RollappPacket_ON_ACK, rollappPortOnHub, rollappChannelOnHub, packet.Sequence)
	if err != nil {
		logger.Error("Failed to get proof height from packet", "err", err)
		return err
	}

	finalized, err := im.CheckIfFinalized(ctx, rollappID, proofHeight)
	if err != nil {
		logger.Error("Failed to check if packet is finalized", "err", err)
		return err
	}

	if finalized {
		logger.Debug("Skipping eIBC transfer OnAcknowledgementPacket as the packet proof height is already finalized")
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}
	// Run the underlying app's OnAcknowledgementPacket callback
	// with cache context to avoid state changes and report the acknowledgement result.
	// Only save the packet if the underlying app's callback succeeds.
	cacheCtx, _ := ctx.CacheContext()
	err = im.IBCModule.OnAcknowledgementPacket(cacheCtx, packet, acknowledgement, relayer)
	if err != nil {
		return err
	}
	// Save the packet data to the store for later processing
	rollappPacket := commontypes.RollappPacket{
		RollappId:       rollappID,
		Packet:          &packet,
		Acknowledgement: acknowledgement,
		Status:          commontypes.Status_PENDING,
		Relayer:         relayer,
		ProofHeight:     proofHeight,
		Type:            commontypes.RollappPacket_ON_ACK,
	}
	im.SetRollappPacket(ctx, rollappPacket)

	logger.Debug("Set rollapp packet",
		"rollappID", rollappPacket.RollappId,
		"proofHeight", rollappPacket.ProofHeight,
		"type", rollappPacket.Type)

	switch ack.Response.(type) {
	// Only if the acknowledgement is an error, we want to create an order
	case *channeltypes.Acknowledgement_Error:
		err = im.eIBCDemandOrderHandler(ctx, rollappPacket, *transferPacketData)
		if err != nil {
			return err
		}
	}

	return im.IBCModule.OnAcknowledgementPacket(cacheCtx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCMiddleware interface
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	wrappedPacket channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	logger := ctx.Logger().With(
		"module", types.ModuleName,
		"packet_source", wrappedPacket.SourcePort,
		"packet_destination", wrappedPacket.DestinationPort,
		"packet_sequence", wrappedPacket.Sequence)

	rollappPortOnHub, rollappChannelOnHub := wrappedPacket.SourcePort, wrappedPacket.SourceChannel

	rollapp, transferPacketData, err := im.rollappKeeper.ExtractRollappAndTransferPacketFromData(
		ctx,
		wrappedPacket.Data,
		rollappPortOnHub,
		rollappChannelOnHub,
	)
	if err != nil {
		logger.Error("Failed to extract rollapp id from channel", "err", err)
		return err
	}

	packet, err := unwrapPacketData(wrappedPacket, transferPacketData)
	if err != nil {
		logger.Error("Failed to marshal transfer packet data", "err", err)
		return err
	}

	if !im.IsRollappsEnabled(ctx) {
		return im.IBCModule.OnTimeoutPacket(ctx, packet, relayer)
	}

	rollappID := rollapp.RollappId

	if rollappID == "" {
		logger.Debug("Skipping IBC transfer OnTimeoutPacket for non-rollapp chain")
		return im.IBCModule.OnTimeoutPacket(ctx, packet, relayer)
	}

	err = im.ValidateRollappId(ctx, rollappID, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		logger.Error("Failed to validate rollappID", "rollappID", rollappID, "err", err)
		return err
	}

	proofHeight, err := im.GetProofHeight(ctx, commontypes.RollappPacket_ON_TIMEOUT, rollappPortOnHub, rollappChannelOnHub, packet.Sequence)
	if err != nil {
		logger.Error("Failed to get proof height from packet", "err", err)
		return err
	}
	finalized, err := im.CheckIfFinalized(ctx, rollappID, proofHeight)
	if err != nil {
		logger.Error("Failed to check if packet is finalized", "err", err)
		return err
	}

	if finalized {
		logger.Debug("Skipping IBC transfer OnTimeoutPacket as the packet proof height is already finalized")
		return im.IBCModule.OnTimeoutPacket(ctx, packet, relayer)
	}

	// Run the underlying app's OnTimeoutPacket callback
	// with cache context to avoid state changes and report the timeout result.
	// Only save the packet if the underlying app's callback succeeds.
	cacheCtx, _ := ctx.CacheContext()
	err = im.IBCModule.OnTimeoutPacket(cacheCtx, packet, relayer)
	if err != nil {
		return err
	}
	// Save the packet data to the store for later processing
	rollappPacket := commontypes.RollappPacket{
		RollappId:   rollappID,
		Packet:      &packet,
		Status:      commontypes.Status_PENDING,
		Relayer:     relayer,
		ProofHeight: proofHeight,
		Type:        commontypes.RollappPacket_ON_TIMEOUT,
	}
	im.SetRollappPacket(ctx, rollappPacket)

	logger.Debug("Set rollapp packet",
		"rollappID", rollappPacket.RollappId,
		"proofHeight", rollappPacket.ProofHeight,
		"type", rollappPacket.Type)

	err = im.eIBCDemandOrderHandler(ctx, rollappPacket, *transferPacketData)
	if err != nil {
		return err
	}

	return nil
}

// GetProofHeight returns the proof height of the packet
func (im IBCMiddleware) GetProofHeight(ctx sdk.Context, packetType commontypes.RollappPacket_Type,
	rollappPortOnHub string, rollappChannelOnHub string, sequence uint64,
) (uint64, error) {
	packetId := commontypes.NewPacketUID(packetType, rollappPortOnHub, rollappChannelOnHub, sequence)
	height, ok := types.FromIBCProofContext(ctx, packetId)
	if ok {
		return height.RevisionHeight, nil
	} else {
		err := errors.New("failed to get proof height from context")
		ctx.Logger().Error(err.Error(), "packetId", packetId)
		return 0, err
	}
}

// CheckIfFinalized checks if the packet is finalized and if so, updates the packet status
func (im IBCMiddleware) CheckIfFinalized(ctx sdk.Context, rollappID string, proofHeight uint64) (bool, error) {
	finalizedHeight, err := im.GetRollappFinalizedHeight(ctx, rollappID)
	if err != nil {
		if errors.Is(err, rollapptypes.ErrNoFinalizedStateYetForRollapp) {
			return false, nil
		}
		return false, err
	}

	return finalizedHeight >= proofHeight, nil
}
