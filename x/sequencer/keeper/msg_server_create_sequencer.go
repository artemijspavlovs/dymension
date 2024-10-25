package keeper

import (
	"context"
	"slices"
	"strconv"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	rollapptypes "github.com/dymensionxyz/dymension/v3/x/rollapp/types"

	"github.com/dymensionxyz/dymension/v3/x/sequencer/types"
)

func (k msgServer) CreateSequencer(goCtx context.Context, msg *types.MsgCreateSequencer) (*types.MsgCreateSequencerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check to see if the rollapp has been registered before
	rollapp, found := k.rollappKeeper.GetRollapp(ctx, msg.RollappId)
	if !found {
		return nil, rollapptypes.ErrRollappNotFound
	}

	// check to see if the seq has been registered before
	if _, err := k.GetRealSequencer(ctx, msg.Creator); err == nil {
		return nil, types.ErrSequencerAlreadyExists
	}

	/*
		If we are awaiting the last block from the proposer we stop new sequencer registrations, because
		we don't want to set a new successor while the last block from the proposer is in flight.
		TODO: possible to simplify?
	*/
	if k.awaitingLastProposerBlock(ctx, msg.RollappId) {
		return nil, types.ErrRegisterSequencerWhileAwaitingLastProposerBlock
	}

	if err := k.sufficientBond(ctx, msg.Bond); err != nil {
		return nil, err
	}

	if err := msg.VMSpecificValidate(rollapp.VmType); err != nil {
		return nil, errorsmod.Wrap(err, "vm specific validate")
	}

	// In case InitialSequencer is set to one or more bech32 addresses, only one of them can be the first to register,
	// and is automatically selected as the first proposer, allowing the Rollapp to be set to 'launched'
	// (provided that all the immutable fields are set in the Rollapp).
	// This limitation prevents scenarios such as:
	// a) any unintended initial seq getting registered before the immutable fields are set in the Rollapp.
	// b) situation when seq "X" is registered prior to the initial seq,
	// after which the initial seq's address is set to seq X's address, effectively preventing:
	// 	1. the initial seq from getting selected as the first proposer,
	// 	2. the rollapp from getting launched again
	// In case the InitialSequencer is set to the "*" wildcard, any seq can be the first to register.
	if !rollapp.Launched {
		isInitialOrAllAllowed := slices.Contains(strings.Split(rollapp.InitialSequencer, ","), msg.Creator) ||
			rollapp.InitialSequencer == "*"
		if !isInitialOrAllAllowed {
			return nil, types.ErrNotInitialSequencer
		}

		// check pre launch time.
		// skipped if no pre launch time is set
		if rollapp.PreLaunchTime != nil && rollapp.PreLaunchTime.After(ctx.BlockTime()) {
			return nil, types.ErrBeforePreLaunchTime
		}

		if err := k.rollappKeeper.SetRollappAsLaunched(ctx, &rollapp); err != nil {
			return nil, err
		}
	}

	seq := k.NewSequencer(ctx, msg.RollappId)
	seq.DymintPubKey = msg.DymintPubKey
	seq.Address = msg.Creator
	seq.Status = types.Bonded
	seq.Metadata = msg.Metadata
	seq.OptedIn = true

	if err := k.sendToModule(ctx, seq, msg.Bond); err != nil {
		return nil, err
	}

	k.SetSequencer(ctx, *seq)

	if err := k.ChooseProposer(ctx, msg.RollappId); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateSequencer,
			sdk.NewAttribute(types.AttributeKeyRollappId, msg.RollappId),
			sdk.NewAttribute(types.AttributeKeySequencer, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyBond, msg.Bond.String()),
			sdk.NewAttribute(types.AttributeKeyProposer, strconv.FormatBool(k.IsProposer(ctx, *seq))),
		),
	)

	return &types.MsgCreateSequencerResponse{}, nil
}
