package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/sdk-utils/utils/uevent"

	"github.com/dymensionxyz/dymension/v3/x/sequencer/types"
)

func (k msgServer) UpdateSequencerInformation(
	goCtx context.Context,
	msg *types.MsgUpdateSequencerInformation,
) (*types.MsgUpdateSequencerInformationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	seq, err := k.GetRealSequencer(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	defer func() {
		k.SetSequencer(ctx, seq)
	}()

	rollapp := k.rollappKeeper.MustGetRollapp(ctx, seq.RollappId)

	if err := msg.VMSpecificValidate(rollapp.VmType); err != nil {
		return nil, err
	}

	seq.Metadata = msg.Metadata

	if err := uevent.EmitTypedEvent(ctx, &seq); err != nil {
		return nil, fmt.Errorf("emit event: %w", err)
	}

	return &types.MsgUpdateSequencerInformationResponse{}, nil
}
