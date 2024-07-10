package v4

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/dymensionxyz/dymension/v3/app/upgrades"
)

const (
	UpgradeName = "v4"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        storetypes.StoreUpgrades{},
}