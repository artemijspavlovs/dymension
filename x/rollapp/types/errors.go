package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/rollapp module sentinel errors
var (
	ErrRollappExists                  = errorsmod.Register(ModuleName, 1000, "rollapp already exists for this rollapp-id; must use new rollapp-id")
	ErrInvalidInitialSequencerAddress = errorsmod.Register(ModuleName, 1001, "empty initial sequencer address")
	ErrInvalidCreatorAddress          = errorsmod.Register(ModuleName, 1002, "invalid creator address")
	ErrInvalidBech32Prefix            = errorsmod.Register(ModuleName, 1003, "invalid Bech32 prefix")
	ErrInvalidNumBlocks               = errorsmod.Register(ModuleName, 1005, "invalid number of blocks")
	ErrInvalidBlockSequence           = errorsmod.Register(ModuleName, 1006, "invalid block sequence")
	ErrUnknownRollappID               = errorsmod.Register(ModuleName, 1007, "rollapp does not exist")
	ErrVersionMismatch                = errorsmod.Register(ModuleName, 1008, "rollapp version mismatch")
	ErrWrongBlockHeight               = errorsmod.Register(ModuleName, 1009, "start-height does not match rollapps state")
	ErrEmptyGenesisChecksum           = errorsmod.Register(ModuleName, 1010, "empty genesis checksum")
	ErrInvalidStateRoot               = errorsmod.Register(ModuleName, 1011, "invalid blocks state root")
	ErrEmptyGenesisURLs               = errorsmod.Register(ModuleName, 1012, "empty genesis URLs")
	ErrFeePayment                     = errorsmod.Register(ModuleName, 1013, "rollapp creation fee payment error")
	ErrStateNotExists                 = errorsmod.Register(ModuleName, 1017, "state of this height doesn't exist")
	ErrInvalidHeight                  = errorsmod.Register(ModuleName, 1018, "invalid rollapp height")
	ErrInvalidRollappID               = errorsmod.Register(ModuleName, 1020, "invalid rollapp-id")
	ErrNoFinalizedStateYetForRollapp  = errorsmod.Register(ModuleName, 1024, "no finalized state yet for rollapp")
	ErrInvalidClientState             = errorsmod.Register(ModuleName, 1025, "invalid client state")
	ErrRollappNotRegistered           = errorsmod.Register(ModuleName, 1035, "rollapp not registered")
	ErrUnknownRequest                 = errorsmod.Register(ModuleName, 1036, "unknown request")
	ErrNotFound                       = errorsmod.Register(ModuleName, 1037, "not found")
	ErrLogic                          = errorsmod.Register(ModuleName, 1038, "internal logic error")
	ErrInvalidAddress                 = errorsmod.Register(ModuleName, 1040, "invalid address")
	ErrInitialSequencerAddressTaken   = errorsmod.Register(ModuleName, 1041, "initial sequencer address already taken")

	/* ------------------------------ fraud related ----------------------------- */
	ErrDisputeAlreadyFinalized = errorsmod.Register(ModuleName, 2000, "disputed height already finalized")
	ErrDisputeAlreadyReverted  = errorsmod.Register(ModuleName, 2001, "disputed height already reverted")
	ErrWrongClientId           = errorsmod.Register(ModuleName, 2002, "client id does not match the rollapp")
	ErrWrongProposerAddr       = errorsmod.Register(ModuleName, 2003, "wrong proposer address")
)
