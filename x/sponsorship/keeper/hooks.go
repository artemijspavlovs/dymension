package keeper

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/dymensionxyz/dymension/v3/x/sponsorship/types"
)

var _ stakingtypes.StakingHooks = Hooks{}

// Hooks wrapper struct for slashing keeper
type Hooks struct {
	k Keeper
}

func (k Keeper) Hooks() Hooks {
	return Hooks{k: k}
}

type processHookResult struct {
	distribution     types.Distribution
	votePruned       bool
	vpDiff, newTotal math.Int
}

func (h Hooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	if !h.k.Voted(ctx, delAddr) {
		return nil
	}
	err := h.afterDelegationModified(ctx, delAddr, valAddr)
	if err != nil {
		return fmt.Errorf("sponsorship: delegator '%s', validator '%s': %w", delAddr, valAddr, err)
	}
	return nil
}

func (h Hooks) afterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	v, found := h.k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return fmt.Errorf("validator not found")
	}

	d, found := h.k.stakingKeeper.GetDelegation(ctx, delAddr, valAddr)
	if !found {
		return fmt.Errorf("delegation not found")
	}

	// Calculate a staking voting power
	stakingVP := v.TokensFromShares(d.GetShares()).Ceil().TruncateInt()

	// Get the current voting power saved in x/sponsorship
	sponsorshipVP, err := h.k.GetDelegatorValidatorPower(ctx, delAddr, valAddr)
	if err != nil {
		return fmt.Errorf("cannot get current voting power: %w", err)
	}

	result, err := h.processHook(ctx, delAddr, valAddr, sponsorshipVP, stakingVP)
	if err != nil {
		return fmt.Errorf("cannot process hook: %w", err)
	}

	err = ctx.EventManager().EmitTypedEvent(&types.EventVotingPowerUpdate{
		Voter:           delAddr.String(),
		Validator:       valAddr.String(),
		Distribution:    result.distribution,
		VotePruned:      result.votePruned,
		NewVotingPower:  result.newTotal,
		VotingPowerDiff: result.vpDiff,
	})
	if err != nil {
		return err
	}

	return nil
}

func (h Hooks) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	if !h.k.Voted(ctx, delAddr) {
		return nil
	}
	err := h.beforeDelegationRemoved(ctx, delAddr, valAddr)
	if err != nil {
		return fmt.Errorf("sponsorship: delegator '%s', validator '%s': %w", delAddr, valAddr, err)
	}
	return nil
}

func (h Hooks) beforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	// Get the current voting power saved in x/sponsorship
	// StakingVP is zero in that case, so Diff == (-1) * SponsorshipVP.
	sponsorshipVP, err := h.k.GetDelegatorValidatorPower(ctx, delAddr, valAddr)
	if err != nil {
		return fmt.Errorf("cannot get current voting power: %w", err)
	}

	result, err := h.processHook(ctx, delAddr, valAddr, sponsorshipVP, math.ZeroInt())
	if err != nil {
		return fmt.Errorf("cannot process hook: %w", err)
	}

	err = ctx.EventManager().EmitTypedEvent(&types.EventVotingPowerUpdate{
		Voter:           delAddr.String(),
		Validator:       valAddr.String(),
		Distribution:    result.distribution,
		VotePruned:      result.votePruned,
		NewVotingPower:  result.newTotal,
		VotingPowerDiff: result.vpDiff,
	})
	if err != nil {
		return err
	}

	return nil
}

func (h Hooks) processHook(
	ctx sdk.Context,
	delAddr sdk.AccAddress,
	valAddr sdk.ValAddress,
	oldVP, newVP math.Int,
) (*processHookResult, error) {
	vote, err := h.k.GetVote(ctx, delAddr)
	if err != nil {
		return nil, fmt.Errorf("could not get vote for delegator '%s': %w", delAddr, err)
	}

	// Calculate the diff: if it's > 0, then the user has bonded. Otherwise, unbonded.
	powerDiff := newVP.Sub(oldVP)

	// Update the current user's voting power
	vote.VotingPower = vote.VotingPower.Add(powerDiff)

	// Validate that the user has min voting power. Revoke the vote if not.
	minVP := h.k.GetParams(ctx).MinVotingPower
	if vote.VotingPower.LT(minVP) {
		distr, errX := h.k.revokeVote(ctx, delAddr, vote)
		if errX != nil {
			return nil, fmt.Errorf("could not revoke vote: %w", errX)
		}
		return &processHookResult{
			distribution: distr,
			votePruned:   true,
			vpDiff:       powerDiff,
			newTotal:     vote.VotingPower,
		}, nil
	}

	// The code below updates the vote

	// Apply the vote weight breakdown to the diff -> get a distribution update in absolute values
	update := types.ApplyWeights(powerDiff, vote.Weights)

	// Update the current distribution
	distr, err := h.k.UpdateDistribution(ctx, update.Merge)
	if err != nil {
		return nil, fmt.Errorf("failed to update distribution: %w", err)
	}

	// Save the updated vote
	err = h.k.SaveVote(ctx, delAddr, vote)
	if err != nil {
		return nil, fmt.Errorf("cannot save vote: %w", err)
	}

	// Delete the record if the new VP is zero. Otherwise, update the existing.
	if newVP.IsZero() {
		// Delete the voting power cast for this validator
		h.k.DeleteDelegatorValidatorPower(ctx, delAddr, valAddr)
	} else {
		// Update the voting power cast for this validator
		err = h.k.SaveDelegatorValidatorPower(ctx, delAddr, valAddr, newVP)
		if err != nil {
			return nil, fmt.Errorf("cannot save voting power: %w", err)
		}
	}

	return &processHookResult{
		distribution: distr,
		votePruned:   false,
		vpDiff:       powerDiff,
		newTotal:     vote.VotingPower,
	}, nil
}

func (h Hooks) AfterValidatorBeginUnbonding(sdk.Context, sdk.ConsAddress, sdk.ValAddress) error {
	return nil
}

func (h Hooks) AfterValidatorBonded(sdk.Context, sdk.ConsAddress, sdk.ValAddress) error { return nil }

func (h Hooks) BeforeValidatorSlashed(sdk.Context, sdk.ValAddress, sdk.Dec) error { return nil }

func (Hooks) AfterValidatorCreated(sdk.Context, sdk.ValAddress) error { return nil }

func (Hooks) BeforeValidatorModified(sdk.Context, sdk.ValAddress) error { return nil }

func (Hooks) AfterValidatorRemoved(sdk.Context, sdk.ConsAddress, sdk.ValAddress) error { return nil }

func (Hooks) BeforeDelegationCreated(sdk.Context, sdk.AccAddress, sdk.ValAddress) error { return nil }

func (Hooks) AfterUnbondingInitiated(sdk.Context, uint64) error { return nil }

func (Hooks) BeforeDelegationSharesModified(sdk.Context, sdk.AccAddress, sdk.ValAddress) error {
	return nil
}
