package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/dymension/v3/x/sequencer/types"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
	"github.com/dymensionxyz/sdk-utils/utils/ucoin"
)

type UnbondBlocker interface {
	// CanUnbond should return a types.UnbondNotAllowed error with a reason, or nil (or another error)
	CanUnbond(ctx sdk.Context, sequencer types.Sequencer) error
}

func (k Keeper) TryUnbond(ctx sdk.Context, seq *types.Sequencer, amt sdk.Coin) error {
	if k.isProposerOrSuccessor(ctx, *seq) {
		return types.ErrUnbondProposerOrSuccessor
	}
	for _, c := range k.unbondBlockers {
		if err := c.CanUnbond(ctx, *seq); err != nil {
			return errorsmod.Wrap(err, "other module")
		}
	}
	bond := seq.TokensCoin()
	minBond := k.GetParams(ctx).MinBond
	maxReduction, _ := bond.SafeSub(minBond)
	isPartial := !amt.IsEqual(bond)
	if isPartial && maxReduction.IsLT(amt) {
		return errorsmod.Wrapf(types.ErrUnbondNotAllowed,
			"attempted reduction: %s, max reduction: %s",
			amt, ucoin.NonNegative(maxReduction),
		)
	}
	if err := k.refund(ctx, seq, amt); err != nil {
		return errorsmod.Wrap(err, "refund")
	}
	if seq.Tokens.IsZero() {
		return errorsmod.Wrap(k.unbond(ctx, seq), "unbond")
	}
	return nil
}

func (k Keeper) unbond(ctx sdk.Context, seq *types.Sequencer) error {
	if k.IsSuccessor(ctx, *seq) {
		return gerrc.ErrInternal.Wrap(`unbond next proposer: it shouldnt be possible because
they cannot do frauds and they cannot unbond gracefully`)
	}
	seq.Status = types.Unbonded
	if k.IsProposer(ctx, *seq) {
		k.SetProposer(ctx, seq.RollappId, types.SentinelSeqAddr)
	}
	return nil
}
