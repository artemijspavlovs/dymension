package bridging_fee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	transferapp "github.com/cosmos/ibc-go/v6/modules/apps/transfer"
	transferkeeper "github.com/cosmos/ibc-go/v6/modules/apps/transfer/keeper"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	delayedaackkeeper "github.com/dymensionxyz/dymension/v3/x/delayedack/keeper"
	"github.com/tendermint/tendermint/libs/log"
)

const (
	ModuleName = "bridging_fee"
)

var _ porttypes.Middleware = &BridgingFeeMiddleware{}

// BridgingFeeMiddleware is responsible for charging a bridging fee on transfers coming from rollapps
// The actual charge happens on the packet finalization
// based on ADR: https://www.notion.so/dymension/ADR-x-Bridging-Fee-Middleware-7ba8c191373f43ce81782fc759913299?pvs=4
type BridgingFeeMiddleware struct {
	transferapp.IBCModule
	porttypes.ICS4Wrapper

	delayedAckKeeper delayedaackkeeper.Keeper
	transferKeeper   transferkeeper.Keeper
	feeModuleAddr    sdk.AccAddress
}

func NewIBCMiddleware(
	transfer transferapp.IBCModule,
	channelKeeper porttypes.ICS4Wrapper,
	keeper delayedaackkeeper.Keeper,
	transferKeeper transferkeeper.Keeper,
	feeModuleAddr sdk.AccAddress,
) *BridgingFeeMiddleware {
	return &BridgingFeeMiddleware{
		IBCModule:        transfer,
		ICS4Wrapper:      channelKeeper,
		delayedAckKeeper: keeper,
		transferKeeper:   transferKeeper,
		feeModuleAddr:    feeModuleAddr,
	}
}

func (im BridgingFeeMiddleware) logger(
	ctx sdk.Context,
	packet channeltypes.Packet,
	method string,
) log.Logger {
	return ctx.Logger().With(
		"module", ModuleName,
		"packet_source_port", packet.SourcePort,
		"packet_destination_port", packet.DestinationPort,
		"packet_sequence", packet.Sequence,
		"method", method,
	)
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

	transfer, err := im.delayedAckKeeper.ExtractValidTransfer(ctx, packet)
	if err != nil {
		logger.Error("Get valid transfer transfer", "err", err)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	if !transfer.HasRollapp() {
		logger.Debug("Skipping IBC transfer OnRecvPacket for non-rollapp chain.")
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	// Use the packet as a basis for a fee transfer
	feeData := transfer
	fee := im.delayedAckKeeper.BridgingFeeFromAmt(ctx, transfer.MustAmountInt())
	feeData.Amount = fee.String()
	feeData.Receiver = im.feeModuleAddr.String()

	// No event emitted, as we called the transfer keeper directly (vs the transfer middleware)
	err = im.transferKeeper.OnRecvPacket(ctx, packet, feeData.FungibleTokenPacketData)
	if err == nil {
		logger.Error("Charge bridging fee.", "err", err)
		// we continue as we don't want the fee charge to fail the transfer in any case
		fee = sdk.ZeroInt()
	} else {
		logger.Debug("Charged bridging fee.", "fee", fee)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				EventTypeBridgingFee,
				sdk.NewAttribute(AttributeKeyFee, fee.String()),
				sdk.NewAttribute(sdk.AttributeKeySender, transfer.Sender),
			),
		)
	}

	// transfer the rest to the original recipient
	transfer.Amount = transfer.MustAmountInt().Sub(fee).String()
	packet.Data = transfer.GetBytes()
	return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
}
