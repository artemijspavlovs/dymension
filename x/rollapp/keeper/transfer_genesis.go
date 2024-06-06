package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dymerror "github.com/dymensionxyz/dymension/v3/x/common/errors"
	"github.com/dymensionxyz/dymension/v3/x/rollapp/types"
)

// VerifyAndRecordGenesisTransfer takes a transfer 'index' from the rollapp sequencer and book keeps it
// If we have previously seen a different n, we reject it, the sequencer is not following protocol.
// If we have previously seen the same IX already, we reject it, as IBC guarantees exactly once delivery, then the sequencer must not be following protocol
// Once we have recorded n indexes, this rollapp can proceed to the next step of the genesis transfer protocol
// Returns the number of transfers recorded so far (including this one)
func (k Keeper) VerifyAndRecordGenesisTransfer(ctx sdk.Context, rollappID string, ix, nTotal uint64) (uint64, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferGenesisMapKeyPrefix))

	nKey := types.TransferGenesisNumKey(rollappID)
	nTotalKey := types.TransferGenesisNumTotalKey(rollappID)
	ixKey := types.TransferGenesisSetMembershipKey(rollappID, ix)

	n := uint64(0)
	/*
		We do all the verification first and only write at the end, to make it easier to reason about partial failures
	*/

	if !!store.Has(nTotalKey) {
		nTotalExistingBz := store.Get(nTotalKey)
		nTotalExisting := sdk.BigEndianToUint64(nTotalExistingBz)
		if nTotal != nTotalExisting {
			return 0, errorsmod.Wrapf(dymerror.ErrProtocolViolation,
				"different num total transfers: got: %d: got previously: %d", nTotal, nTotalExisting)
		}
		nBz := store.Get(nKey)
		n = sdk.BigEndianToUint64(nBz)
	}
	if !(0 <= ix && ix < nTotal) {
		return 0, errorsmod.Wrapf(dymerror.ErrProtocolViolation,
			"ix must be less than nTotal: ix: %d: nTotal: %d", ix, nTotal)
	}
	if store.Has(ixKey) {
		return 0, errorsmod.Wrapf(dymerror.ErrProtocolViolation,
			"already received genesis transfer: ix: %d", ix)
	}

	n++
	store.Set(nTotalKey, sdk.Uint64ToBigEndian(nTotal))
	store.Set(nKey, sdk.Uint64ToBigEndian(n))
	store.Set(ixKey, []byte{})
	return n, nil
}

func (k Keeper) enableTransfers(ctx sdk.Context, rollappID string) {
	ra := k.MustGetRollapp(ctx, rollappID)
	ra.GenesisState.TransfersEnabled = true
	k.SetRollapp(ctx, ra)
}

func (k Keeper) FinalizeGenesisTransferDisputeWindows(ctx sdk.Context) error { // TODO: needs to be public?

	h := uint64(ctx.BlockHeight())
	if h <= k.DisputePeriodTransferGenesisInBlocks(ctx) { // TODO: check <= is ok
		return nil
	}

	toFinalize := k.GetTransferGenesisFinalizationQueue(ctx, h-k.DisputePeriodTransferGenesisInBlocks(ctx)+1)
	for _, rollapps := range toFinalize {
		for _, raID := range rollapps.Rollapps {
			ra := k.MustGetRollapp(ctx, raID)
			ra.GenesisState.TransfersEnabled = true
			k.SetRollapp(ctx, ra)
			k.DelGenesisTransfers(ctx, raID)
			ctx.EventManager().EmitEvent(transfersEnabledEvent(raID))
			k.Logger(ctx).Info("Genesis transfer window finalized. Transfers enabled.", "rollapp", raID)
		}
		k.DelTransferGenesisFinalizations(ctx, rollapps.CreationHeight)
	}

	return nil
}

func transfersEnabledEvent(raID string) sdk.Event {
	return sdk.NewEvent(types.EventTypeTransferGenesisH2,
		sdk.NewAttribute(types.AttributeKeyRollappId, raID),
	)
}

// TODO: expensive? explain
func (k Keeper) GetAllGenesisTransfers(ctx sdk.Context) []types.GenesisTransfers { // TODO: needs to be public?
	var ret []types.GenesisTransfers

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferGenesisMapKeyPrefix))

	for _, ra := range k.GetAllRollapps(ctx) {

		raID := ra.RollappId
		nTotalKey := types.TransferGenesisNumTotalKey(raID)
		nTotalBz := store.Get(nTotalKey)
		nTotal := sdk.BigEndianToUint64(nTotalBz)
		nKey := types.TransferGenesisNumKey(raID)
		nBz := store.Get(nKey)
		n := sdk.BigEndianToUint64(nBz)
		x := types.GenesisTransfers{
			RollappID:   raID,
			NumTotal:    nTotal,
			NumReceived: n,
		}
		for ix := range nTotal {
			ixKey := types.TransferGenesisSetMembershipKey(raID, ix)
			if store.Has(ixKey) {
				x.Received = append(x.Received, ix)
			}
		}

		ret = append(ret, x)
	}

	// TODO: impl
	return ret
}

// DelGenesisTransfers deletes bookkeeping for one rollapp, it's not needed for correctness, but it's good to prune state
// TODO: need to justify correct pruning property for this and the other state, in all possibilities (e.g. fraud recovery)
func (k Keeper) DelGenesisTransfers(ctx sdk.Context, raID string) { // TODO: needs to be public?

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferGenesisMapKeyPrefix))

	nTotalKey := types.TransferGenesisNumTotalKey(raID)
	nTotalBz := store.Get(nTotalKey)
	nTotal := sdk.BigEndianToUint64(nTotalBz)
	nKey := types.TransferGenesisNumKey(raID)
	store.Delete(nTotalBz)
	store.Delete(nKey)

	for ix := range nTotal {
		store.Delete(types.TransferGenesisSetMembershipKey(raID, ix))
	}
}

func (k Keeper) AddTransferGenesisFinalization(ctx sdk.Context, rollappID string) {
	h := uint64(ctx.BlockHeight())
	fs := k.GetTransferGenesisFinalizations(ctx, h)
	fs.Rollapps = append(fs.Rollapps, rollappID) // we assume it's not there already
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferGenesisQueueKeyPrefix))
	b := k.cdc.MustMarshal(&fs)
	store.Set(types.TransferGenesisFinalizationsKey(h), b)
}

func (k Keeper) GetTransferGenesisFinalizations(
	ctx sdk.Context,
	creationHeight uint64,
) types.BlockHeightToTransferGenesisFinalizations {
	var ret types.BlockHeightToTransferGenesisFinalizations
	ret.Rollapps = make([]string, 0)
	ret.CreationHeight = creationHeight

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferGenesisQueueKeyPrefix))
	b := store.Get(types.TransferGenesisFinalizationsKey(creationHeight))
	if b == nil {
		return ret
	}

	k.cdc.MustUnmarshal(b, &ret)
	return ret
}

func (k Keeper) DelTransferGenesisFinalizations(
	ctx sdk.Context,
	creationHeight uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferGenesisQueueKeyPrefix))
	store.Delete(types.TransferGenesisFinalizationsKey(creationHeight))
}

// GetTransferGenesisFinalizationQueue returns the queue up to but not including height
// Passing height = 0 returns all
func (k Keeper) GetTransferGenesisFinalizationQueue(ctx sdk.Context, creationHeight uint64) []types.BlockHeightToTransferGenesisFinalizations {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferGenesisQueueKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close() // nolint: errcheck

	var ret []types.BlockHeightToTransferGenesisFinalizations
	for ; iterator.Valid(); iterator.Next() {
		var x types.BlockHeightToTransferGenesisFinalizations
		k.cdc.MustUnmarshal(iterator.Value(), &x)
		if creationHeight != 0 && creationHeight <= x.CreationHeight {
			break
		}
		ret = append(ret, x)
	}

	return ret
}
