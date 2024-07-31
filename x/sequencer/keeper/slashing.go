package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/dymension/v3/x/sequencer/types"
)

// Slashing slashes the sequencer for misbehaviour
// Slashing can occur on both Bonded and Unbonding sequencers
func (k Keeper) Slashing(ctx sdk.Context, seqAddr string) error {
	seq, found := k.GetSequencer(ctx, seqAddr)
	if !found {
		return types.ErrUnknownSequencer
	}

	if seq.Status == types.Unbonded {
		return errorsmod.Wrap(
			types.ErrInvalidSequencerStatus,
			"can't slash unbonded sequencer",
		)
	}

	seqTokens := seq.Tokens
	if !seqTokens.Empty() {
		err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, seqTokens)
		if err != nil {
			return err
		}
	} else {
		k.Logger(ctx).Error("sequencer has no tokens to slash", "sequencer", seq.SequencerAddress)
	}
	seq.Tokens = sdk.Coins{}
	oldStatus := seq.Status

	// remove from queue if unbonding
	if oldStatus == types.Unbonding {
		k.removeUnbondingSequencer(ctx, seq)
	}
	// remove from notice period queue if needed
	if seq.IsNoticePeriodInProgress() {
		k.removeNoticePeriodSequencer(ctx, seq)
	}
	// if the slashed sequencer is the proposer, remove it
	// the caller should rotate the proposer
	if k.isProposer(ctx, seq.RollappId, seqAddr) {
		k.removeProposer(ctx, seq.RollappId)
	}

	// if we slash the next proposer, we're in the middle of rotation
	// instead of removing the next proposer, we set it to empty
	if k.isNextProposer(ctx, seq.RollappId, seqAddr) {
		k.setNextProposer(ctx, seq.RollappId, "")
	}

	// set the status to unbonded
	seq.Status = types.Unbonded
	seq.Jailed = true

	seq.UnbondRequestHeight = ctx.BlockHeight()
	seq.UnbondTime = ctx.BlockTime()
	k.UpdateSequencer(ctx, seq, oldStatus)

	// emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSlashed,
			sdk.NewAttribute(types.AttributeKeySequencer, seqAddr),
			sdk.NewAttribute(types.AttributeKeyBond, seqTokens.String()),
		),
	)

	return nil
}
