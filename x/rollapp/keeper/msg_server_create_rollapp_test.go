package keeper_test

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/rand"

	"github.com/dymensionxyz/dymension/v3/testutil/sample"
	"github.com/dymensionxyz/dymension/v3/x/rollapp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *RollappTestSuite) TestCreateRollapp() {
	suite.SetupTest()
	suite.createRollapp(nil)
}

func (suite *RollappTestSuite) TestCreateRollappUnauthorizedRollappCreator() {
	suite.SetupTest()
	suite.createRollappWithCreatorAndVerify(types.ErrFeePayment, bob) // bob is broke
}

func (suite *RollappTestSuite) TestCreateRollappAlreadyExists() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	// rollapp is the rollapp to create
	rollapp := types.MsgCreateRollapp{
		Creator:                 alice,
		RollappId:               "rollapp1",
		InitialSequencerAddress: sample.AccAddress(),
		Bech32Prefix:            bech32Prefix,
		GenesisInfo:             genesisInfo,
	}
	_, err := suite.msgServer.CreateRollapp(goCtx, &rollapp)
	suite.Require().Nil(err)

	_, err = suite.msgServer.CreateRollapp(goCtx, &rollapp)
	suite.ErrorIs(err, types.ErrRollappExists)
}

func (suite *RollappTestSuite) TestCreateRollappSequencerAlreadyExists() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	seqAddr := sample.AccAddress()

	rollapp := types.MsgCreateRollapp{
		Creator:                 alice,
		RollappId:               "rollapp1",
		InitialSequencerAddress: seqAddr,
		Bech32Prefix:            bech32Prefix,
		GenesisInfo:             genesisInfo,
	}
	_, err := suite.msgServer.CreateRollapp(goCtx, &rollapp)
	suite.Require().Nil(err)

	rollapp = types.MsgCreateRollapp{
		Creator:                 alice,
		RollappId:               "rollapp2",
		InitialSequencerAddress: seqAddr,
		Bech32Prefix:            bech32Prefix,
		GenesisInfo:             genesisInfo,
	}
	_, err = suite.msgServer.CreateRollapp(goCtx, &rollapp)
	suite.ErrorIs(err, types.ErrInitialSequencerAddressTaken)
}

func (suite *RollappTestSuite) TestCreateRollappWhenDisabled() {
	suite.SetupTest()

	suite.createRollappAndVerify(nil)
	params := suite.App.RollappKeeper.GetParams(suite.Ctx)
	params.RollappsEnabled = false

	suite.App.RollappKeeper.SetParams(suite.Ctx, params)
	suite.createRollappAndVerify(types.ErrRollappsDisabled)
}

func (suite *RollappTestSuite) TestCreateRollappId() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	tests := []struct {
		name      string
		rollappId string
		eip       bool
		valid     bool
	}{
		{
			name:      "default is valid",
			rollappId: "rollapp_1234-1",
			eip:       true,
			valid:     true,
		},
		{
			name:      "valid non-eip155",
			rollappId: "testChain3",
			eip:       false,
			valid:     true,
		},
		{
			name:      "too long id",
			rollappId: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			eip:       false,
			valid:     false,
		},
		{
			name:      "wrong EIP155",
			rollappId: "rollapp_ea2413-1",
			eip:       false,
			valid:     true,
		},
		{
			name:      "no EIP155 with revision",
			rollappId: "rollapp-1",
			eip:       false,
			valid:     true,
		},
		{
			name:      "starts with dash",
			rollappId: "-1234",
			eip:       false,
			valid:     false,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			rollapp := types.MsgCreateRollapp{
				Creator:                 alice,
				RollappId:               test.rollappId,
				InitialSequencerAddress: sample.AccAddress(),
				Bech32Prefix:            bech32Prefix,
				GenesisInfo:             genesisInfo,
			}

			_, err := suite.msgServer.CreateRollapp(goCtx, &rollapp)
			if test.valid {
				suite.Require().NoError(err)
				id, err := types.NewChainID(test.rollappId)
				suite.Require().NoError(err)
				if test.eip {
					suite.Require().True(id.IsEIP155())
				} else {
					suite.Require().False(id.IsEIP155())
				}
			} else {
				suite.Require().ErrorIs(err, types.ErrInvalidRollappID)
			}
		})
	}
}

func (suite *RollappTestSuite) TestCreateRollappIdRevisionNumber() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	tests := []struct {
		name      string
		rollappId string
		revision  uint64
		valid     bool
	}{
		{
			name:      "revision set with eip155",
			rollappId: "rollapp_1234-1",
			revision:  1,
			valid:     true,
		},
		{
			name:      "revision set without eip155",
			rollappId: "rollapp-3",
			revision:  3,
			valid:     true,
		},
		{
			name:      "revision not set",
			rollappId: "rollapp",
			revision:  0,
			valid:     true,
		},
		{
			name:      "invalid revision",
			rollappId: "rollapp-1-1",
			revision:  0,
			valid:     false,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			rollapp := types.MsgCreateRollapp{
				Creator:                 alice,
				RollappId:               test.rollappId,
				InitialSequencerAddress: sample.AccAddress(),
				Bech32Prefix:            bech32Prefix,
				GenesisInfo:             genesisInfo,
			}

			_, err := suite.msgServer.CreateRollapp(goCtx, &rollapp)

			if test.valid {
				suite.Require().NoError(err)
				id, err := types.NewChainID(test.rollappId)
				suite.Require().NoError(err)
				suite.Require().Equal(test.revision, id.GetRevisionNumber())

			} else {
				suite.Require().ErrorIs(err, types.ErrInvalidRollappID)
			}
		})
	}
}

func (suite *RollappTestSuite) TestForkChainId() {
	tests := []struct {
		name         string
		rollappId    string
		newRollappId string
		valid        bool
	}{
		{
			name:         "valid eip155 id",
			rollappId:    "rollapp_1234-1",
			newRollappId: "rollapp_1234-2",
			valid:        true,
		},
		{
			name:         "valid non-eip155 id",
			rollappId:    "rollapp",
			newRollappId: "rollapp-2",
			valid:        true,
		},
		{
			name:         "non-valid eip155 id",
			rollappId:    "rollapp_1234-1",
			newRollappId: "rollapp_1234-5",
			valid:        false,
		},
		{
			name:         "same eip155 but different name",
			rollappId:    "rollapp_1234-1",
			newRollappId: "rollapy_1234-2",
			valid:        false,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()
			goCtx := sdk.WrapSDKContext(suite.Ctx)
			rollappMsg := types.MsgCreateRollapp{
				Creator:                 alice,
				RollappId:               test.rollappId,
				InitialSequencerAddress: sample.AccAddress(),
				Bech32Prefix:            bech32Prefix,
				GenesisInfo:             genesisInfo,
			}

			_, err := suite.msgServer.CreateRollapp(goCtx, &rollappMsg)
			suite.Require().NoError(err)
			rollapp, found := suite.App.RollappKeeper.GetRollapp(suite.Ctx, rollappMsg.RollappId)
			suite.Require().True(found)
			rollapp.Frozen = true
			suite.App.RollappKeeper.SetRollapp(suite.Ctx, rollapp)

			rollappMsg2 := types.MsgCreateRollapp{
				Creator:                 alice,
				RollappId:               test.newRollappId,
				InitialSequencerAddress: sample.AccAddress(),
				Bech32Prefix:            bech32Prefix,
				GenesisInfo:             genesisInfo,
			}
			_, err = suite.msgServer.CreateRollapp(goCtx, &rollappMsg2)
			if test.valid {
				suite.Require().NoError(err)
				_, found = suite.App.RollappKeeper.GetRollapp(suite.Ctx, rollappMsg2.RollappId)
				suite.Require().True(found)
			} else {
				suite.Require().ErrorIs(err, types.ErrInvalidRollappID)
			}
		})
	}
}

func (suite *RollappTestSuite) TestOverwriteEIP155Key() {
	tests := []struct {
		name         string
		rollappId    string
		badRollappId string
	}{
		{
			name:         "extra whitespace id",
			rollappId:    "rollapp_1234-1",
			badRollappId: "rollapp_1234-1  ",
		},
		{
			name:         "same EIP ID",
			rollappId:    "rollapp_1234-1",
			badRollappId: "dummy_1234-1",
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()
			goCtx := sdk.WrapSDKContext(suite.Ctx)
			rollapp := types.MsgCreateRollapp{
				Creator:                 alice,
				RollappId:               test.rollappId,
				InitialSequencerAddress: sample.AccAddress(),
				Bech32Prefix:            bech32Prefix,
				GenesisInfo:             genesisInfo,
			}
			_, err := suite.msgServer.CreateRollapp(goCtx, &rollapp)
			suite.Require().NoError(err)

			// get eip155 key
			id, err := types.NewChainID(test.rollappId)

			suite.Require().NoError(err)
			suite.Require().NotEqual(0, id.GetEIP155ID())
			eip155key := id.GetEIP155ID()
			// eip155 key registers to correct roll app
			rollAppfromEip1155, found := suite.App.RollappKeeper.GetRollappByEIP155(suite.Ctx, eip155key)
			suite.Require().True(found)
			suite.Require().Equal(rollAppfromEip1155.RollappId, rollapp.RollappId)
			// create bad rollapp
			badrollapp := types.MsgCreateRollapp{
				Creator:                 alice,
				RollappId:               test.badRollappId,
				InitialSequencerAddress: sample.AccAddress(),
				Bech32Prefix:            bech32Prefix,
				GenesisInfo:             genesisInfo,
			}
			_, err = suite.msgServer.CreateRollapp(goCtx, &badrollapp)
			// it should not be possible to register rollapp name with extra space
			suite.Require().ErrorIs(err, types.ErrRollappExists)
		})
	}
}

var (
	genesisInfo = &types.GenesisInfo{
		GenesisUrls:     []string{"https://example.com/genesis.json"},
		GenesisChecksum: "1234567890abcdef",
	}
)

func (suite *RollappTestSuite) createRollapp(expectedErr error) {
	// rollappsExpect is the expected result of query all
	var rollappsExpect []*types.RollappSummary

	// test 10 rollapp creations
	for i := 0; i < 10; i++ {
		res := suite.createRollappAndVerify(expectedErr)
		rollappsExpect = append(rollappsExpect, &res)
	}

	// verify that query all contains all the rollapps that were created
	rollappsRes, totalRes := getAll(suite)
	if expectedErr != nil {
		suite.Require().EqualValues(totalRes, 0)
		return
	} else {
		suite.Require().EqualValues(totalRes, 10)
		verifyAll(suite, rollappsExpect, rollappsRes)
	}
}

func (suite *RollappTestSuite) createRollappAndVerify(expectedErr error) types.RollappSummary {
	return suite.createRollappWithCreatorAndVerify(expectedErr, alice)
}

func (suite *RollappTestSuite) createRollappWithCreatorAndVerify(expectedErr error, creator string) types.RollappSummary {
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	// generate sequencer address
	address := sample.AccAddress()
	// rollapp is the rollapp to create
	rollapp := types.MsgCreateRollapp{
		Creator:                 creator,
		RollappId:               fmt.Sprintf("%s%d", "rollapp", rand.Int63()), //nolint:gosec // this is for a test
		InitialSequencerAddress: address,
		Bech32Prefix:            bech32Prefix,
		GenesisInfo:             genesisInfo,
	}
	// rollappExpect is the expected result of creating rollapp
	rollappExpect := types.Rollapp{
		RollappId:               rollapp.GetRollappId(),
		Creator:                 rollapp.GetCreator(),
		Version:                 0,
		InitialSequencerAddress: rollapp.GetInitialSequencerAddress(),
		Bech32Prefix:            rollapp.GetBech32Prefix(),
		GenesisInfo:             rollapp.GetGenesisInfo(),
	}
	// create rollapp
	createResponse, err := suite.msgServer.CreateRollapp(goCtx, &rollapp)
	if expectedErr != nil {
		suite.ErrorIs(err, expectedErr)
		return types.RollappSummary{}
	}
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgCreateRollappResponse{}, *createResponse)

	// query the specific rollapp
	queryResponse, err := suite.queryClient.Rollapp(goCtx, &types.QueryGetRollappRequest{
		RollappId: rollapp.GetRollappId(),
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(&rollappExpect, &queryResponse.Rollapp)

	rollappSummaryExpect := types.RollappSummary{
		RollappId: rollappExpect.RollappId,
	}
	return rollappSummaryExpect
}
