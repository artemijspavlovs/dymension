package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
)

var (
	ErrRollappExists                 = errorsmod.Register(ModuleName, 1000, "rollapp already exists")
	ErrInvalidInitialSequencer       = errorsmod.Register(ModuleName, 1001, "empty initial sequencer")
	ErrInvalidCreatorAddress         = errorsmod.Register(ModuleName, 1002, "invalid creator address")
	ErrRollappFrozen                 = errorsmod.Register(ModuleName, 1004, "rollapp is frozen")
	ErrInvalidNumBlocks              = errorsmod.Register(ModuleName, 1005, "invalid number of blocks")
	ErrInvalidBlockSequence          = errorsmod.Register(ModuleName, 1006, "invalid block sequence")
	ErrUnknownRollappID              = errorsmod.Register(ModuleName, 1007, "rollapp does not exist")
	ErrWrongBlockHeight              = errorsmod.Register(ModuleName, 1009, "start-height does not match rollapps state")
	ErrInvalidGenesisChecksum        = errorsmod.Register(ModuleName, 1010, "invalid genesis checksum")
	ErrInvalidStateRoot              = errorsmod.Register(ModuleName, 1011, "invalid blocks state root")
	ErrAppRegistrationFeePayment     = errorsmod.Register(ModuleName, 1013, "app registration fee payment error")
	ErrStateNotExists                = gerrc.ErrNotFound.Wrap("state of this height doesn't exist")
	ErrInvalidHeight                 = errorsmod.Register(ModuleName, 1018, "invalid rollapp height")
	ErrInvalidRollappID              = errorsmod.Register(ModuleName, 1020, "invalid rollapp-id")
	ErrNoFinalizedStateYetForRollapp = errorsmod.Register(ModuleName, 1024, "no finalized state yet for rollapp")
	ErrInvalidClientState            = errorsmod.Register(ModuleName, 1025, "invalid client state")
	ErrRollappNotRegistered          = errorsmod.Register(ModuleName, 1035, "rollapp not registered")
	ErrUnknownRequest                = errorsmod.Register(ModuleName, 1036, "unknown request")
	ErrNotFound                      = errorsmod.Register(ModuleName, 1037, "not found")
	ErrLogic                         = errorsmod.Register(ModuleName, 1038, "internal logic error")
	ErrInvalidAddress                = errorsmod.Register(ModuleName, 1040, "invalid address")

	ErrInvalidAlias                      = errorsmod.Wrap(gerrc.ErrInvalidArgument, "alias")
	ErrInvalidURL                        = errorsmod.Wrap(gerrc.ErrInvalidArgument, "invalid url")
	ErrInvalidDescription                = errorsmod.Wrap(gerrc.ErrInvalidArgument, "description")
	ErrInvalidMetadata                   = errorsmod.Wrap(gerrc.ErrInvalidArgument, "metadata")
	ErrImmutableFieldUpdateAfterLaunched = errorsmod.Wrap(gerrc.ErrInvalidArgument, "update immutable field after rollapp launched")
	ErrGenesisInfoSealed                 = errorsmod.Wrap(gerrc.ErrInvalidArgument, "update genesis info sealed")
	ErrUnauthorizedSigner                = errorsmod.Wrap(gerrc.ErrPermissionDenied, "unauthorized signer")
	ErrSameOwner                         = errorsmod.Wrap(gerrc.ErrInvalidArgument, "same owner")
	ErrInvalidRequest                    = errorsmod.Wrap(gerrc.ErrInvalidArgument, "invalid request")
	ErrInvalidVMType                     = errorsmod.Wrap(gerrc.ErrInvalidArgument, "invalid vm type")
	ErrInvalidAppID                      = errorsmod.Wrap(gerrc.ErrInvalidArgument, "app id")
	ErrInvalidAppName                    = errorsmod.Wrap(gerrc.ErrInvalidArgument, "app name")
	ErrInvalidAppImage                   = errorsmod.Wrap(gerrc.ErrInvalidArgument, "app image path")
	ErrInvalidBlockDescriptorTimestamp   = errorsmod.Wrap(gerrc.ErrInvalidArgument, "invalid block descriptor timestamp")
	ErrInvalidNativeDenom                = errorsmod.Wrap(gerrc.ErrInvalidArgument, "invalid native denom")
	ErrInvalidFeeDenom                   = errorsmod.Wrap(gerrc.ErrInvalidArgument, "invalid fee denom")
	ErrInvalidInitialSupply              = errorsmod.Wrap(gerrc.ErrInvalidArgument, "invalid initial supply")
	ErrInvalidBech32Prefix               = errorsmod.Wrap(gerrc.ErrInvalidArgument, "invalid bech32 prefix")
	ErrRollappNotFound                   = errorsmod.Wrap(gerrc.ErrNotFound, "rollapp")

	/* ------------------------------ fraud related ----------------------------- */
	ErrDisputeAlreadyFinalized = errorsmod.Register(ModuleName, 2000, "disputed height already finalized")
	ErrDisputeAlreadyReverted  = errorsmod.Register(ModuleName, 2001, "disputed height already reverted")
	ErrWrongClientId           = errorsmod.Register(ModuleName, 2002, "client id does not match the rollapp")
	ErrWrongProposerAddr       = errorsmod.Register(ModuleName, 2003, "wrong proposer address")
	ErrInvalidDRSVersion       = errorsmod.Register(ModuleName, 2004, "wrong DRS version")
)
