package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"github.com/cometbft/cometbft/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dymensionxyz/dymension/v3/x/sequencer/types"
)

type Keeper struct {
	authority string // authority is the x/gov module account

	cdc            codec.BinaryCodec
	storeKey       storetypes.StoreKey
	bankKeeper     types.BankKeeper
	rollappKeeper  types.RollappKeeper
	unbondBlockers []UnbondBlocker
	hooks          types.Hooks

	dymintProposerAddrToAccAddr collections.Map[[]byte, string]
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	bankKeeper types.BankKeeper,
	rollappKeeper types.RollappKeeper,
	authority string,
) *Keeper {
	_, err := sdk.AccAddressFromBech32(authority)
	if err != nil {
		panic(fmt.Errorf("invalid x/sequencer authority address: %w", err))
	}

	return &Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		bankKeeper:     bankKeeper,
		rollappKeeper:  rollappKeeper,
		authority:      authority,
		unbondBlockers: []UnbondBlocker{},
		hooks:          types.NoOpHooks{},
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) AddUnbondBlockers(ubs ...UnbondBlocker) {
	k.unbondBlockers = append(k.unbondBlockers, ubs...)
}

func (k *Keeper) SetHooks(h types.Hooks) {
	k.hooks = h
}
