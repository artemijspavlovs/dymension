package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	rollapptypes "github.com/dymensionxyz/dymension/v3/x/rollapp/types"
	"github.com/dymensionxyz/dymension/v3/x/sequencer/types"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
)

var _ rollapptypes.RollappHooks = rollappHook{}

// Hooks wrapper struct for rollapp keeper.
type rollappHook struct {
	k Keeper
}

// RollappHooks returns the wrapper struct.
func (k Keeper) RollappHooks() rollapptypes.RollappHooks {
	return rollappHook{
		k,
	}
}

// BeforeUpdateState checks various conditions before updating the state.
// It verifies if the sequencer has been registered, if the rollappId matches the one of the sequencer,
// if there is a proposer for the given rollappId, and if the sequencer is the active one.
// If the lastStateUpdateBySequencer flag is true, it also checks if the rollappId is rotating and
// performs a rotation of the proposer.
// Returns an error if any of the checks fail, otherwise returns nil.
func (hook rollappHook) BeforeUpdateState(ctx sdk.Context, seqAddr, rollappId string, lastStateUpdateBySequencer bool) error {
	// check to see if the sequencer has been registered before
	sequencer, found := hook.k.GetSequencer(ctx, seqAddr)
	if !found {
		return types.ErrUnknownSequencer
	}

	// check to see if the rollappId matches the one of the sequencer
	if sequencer.RollappId != rollappId {
		return types.ErrSequencerRollappMismatch
	}

	seq, ok := hook.k.GetProposer(ctx, rollappId)
	if !ok {
		return errors.Join(gerrc.ErrNotFound, types.ErrNoProposer)
	}
	if sequencer.SequencerAddress != seq.SequencerAddress {
		return types.ErrNotActiveSequencer
	}

	if lastStateUpdateBySequencer {
		if !hook.k.IsRotating(ctx, rollappId) {
			return types.ErrInvalidRequest
		}
		// TODO: the hub should probably validate the lastBlock in the lastBatch,
		// to make sure the sequencer is passing the correct nextSequencer on the L2

		hook.k.RotateProposer(ctx, rollappId)
	}

	return nil
}

func (hook rollappHook) AfterStateFinalized(ctx sdk.Context, rollappID string, stateInfo *rollapptypes.StateInfo) error {
	return nil
}

// FraudSubmitted implements the RollappHooks interface
// It slashes the sequencer and unbonds all other bonded sequencers
func (hook rollappHook) FraudSubmitted(ctx sdk.Context, rollappID string, height uint64, seqAddr string) error {
	err := hook.k.Slashing(ctx, seqAddr)
	if err != nil {
		return err
	}

	// unbond all other sequencers
	bonded := hook.k.GetSequencersByRollappByStatus(ctx, rollappID, types.Bonded)
	unbonding := hook.k.GetSequencersByRollappByStatus(ctx, rollappID, types.Unbonding)
	for _, sequencer := range append(bonded, unbonding...) {
		err := hook.k.unbondSequencer(ctx, sequencer.SequencerAddress)
		if err != nil {
			return err
		}
	}

	// clear the proposer and next proposer
	hook.k.removeProposer(ctx, rollappID)
	hook.k.removeNextProposer(ctx, rollappID)

	return nil
}

// RollappCreated implements types.RollappHooks.
func (hook rollappHook) RollappCreated(ctx sdk.Context, rollappID string) error {
	return nil
}
