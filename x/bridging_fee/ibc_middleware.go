package bridging_fee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfer "github.com/cosmos/ibc-go/v6/modules/apps/transfer"
	transferkeeper "github.com/cosmos/ibc-go/v6/modules/apps/transfer/keeper"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	delayedaackkeeper "github.com/dymensionxyz/dymension/v3/x/delayedack/keeper"
)

const (
	ModuleName = "bridging_fee"
)

var _ porttypes.Middleware = &BridgingFeeMiddleware{}

// BridgingFeeMiddleware implements the ICS26 callbacks
// The middleware is responsible for charging a bridging fee on transfers coming from rollapps
// The actual charge happens on the packet finalization
// based on ADR: https://www.notion.so/dymension/ADR-x-Bridging-Fee-Middleware-7ba8c191373f43ce81782fc759913299?pvs=4
type BridgingFeeMiddleware struct {
	transfer.IBCModule
	porttypes.ICS4Wrapper

	delayedAckKeeper delayedaackkeeper.Keeper
	transferKeeper   transferkeeper.Keeper
	feeModuleAddr    sdk.AccAddress
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application
func NewIBCMiddleware(transfer transfer.IBCModule, channelKeeper porttypes.ICS4Wrapper, keeper delayedaackkeeper.Keeper, transferKeeper transferkeeper.Keeper, feeModuleAddr sdk.AccAddress) *BridgingFeeMiddleware {
	return &BridgingFeeMiddleware{
		IBCModule:        transfer,
		ICS4Wrapper:      channelKeeper,
		delayedAckKeeper: keeper,
		transferKeeper:   transferKeeper,
		feeModuleAddr:    feeModuleAddr,
	}
}

// GetFeeRecipient returns the address that will receive the bridging fee
func (im BridgingFeeMiddleware) GetFeeRecipient() sdk.AccAddress {
	return im.feeModuleAddr
}

func (im *BridgingFeeMiddleware) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) exported.Acknowledgement {
	if !im.delayedAckKeeper.IsRollappsEnabled(ctx) {
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}
	logger := ctx.Logger().With(
		"module", ModuleName,
		"packet_source", packet.SourcePort,
		"packet_destination", packet.DestinationPort,
		"packet_sequence", packet.Sequence)

	rollappPortOnHub, rollappChannelOnHub := packet.DestinationPort, packet.DestinationChannel

	data, err := im.delayedAckKeeper.GetRollappAndTransferDataFromPacket(ctx, packet, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		logger.Error("Get transfer data from packet.", "err", err)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	if data.RollappID == "" {
		logger.Debug("Skipping IBC transfer OnRecvPacket for non-rollapp chain.")
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	if err := data.ValidateBasic(); err != nil { // TODO: double check ok, need to wrap?
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// get fee
	fee := im.delayedAckKeeper.BridgingFeeFromAmt(ctx, data.MustAmountInt())

	// update packet data for the fee charge
	feePacket := *transferPacketData
	feePacket.Amount = fee.String()
	feePacket.Receiver = im.GetFeeRecipient().String()

	// No event emitted, as we called the transfer keeper directly (vs the transfer middleware)
	err = im.transferKeeper.OnRecvPacket(ctx, packet, feePacket)
	if err != nil {
		logger.Error("Failed to charge bridging fee", "err", err)
		// we continue as we don't want the fee charge to fail the transfer in any case
		fee = sdk.ZeroInt()
	} else {
		logger.Debug("Charged bridging fee", "fee", fee)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				EventTypeBridgingFee,
				sdk.NewAttribute(AttributeKeyFee, fee.String()),
				sdk.NewAttribute(sdk.AttributeKeySender, transferPacketData.Sender),
			),
		)
	}

	// transfer the remaining amount
	transferPacketData.Amount = transferAmount.Sub(fee).String()
	packet.Data = transferPacketData.GetBytes()
	return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
}
