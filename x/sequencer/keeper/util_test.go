package keeper_test

import (
	"reflect"
	"strconv"
	"testing"

	cometbftproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankutil "github.com/cosmos/cosmos-sdk/x/bank/testutil"
	"github.com/dymensionxyz/dymension/v3/app/apptesting"
	rollappkeeper "github.com/dymensionxyz/dymension/v3/x/rollapp/keeper"
	rollapptypes "github.com/dymensionxyz/dymension/v3/x/rollapp/types"
	"github.com/dymensionxyz/sdk-utils/utils/urand"
	"github.com/stretchr/testify/suite"

	"github.com/dymensionxyz/dymension/v3/x/sequencer/keeper"
	"github.com/dymensionxyz/dymension/v3/x/sequencer/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/dymensionxyz/dymension/v3/testutil/sample"
)

var (
	bond = types.DefaultParams().MinBond
	kick = types.DefaultParams().KickThreshold
	pks  = []cryptotypes.PubKey{
		randPK(),
		randPK(),
		randPK(),
		randPK(),
		randPK(),
		randPK(),
		randPK(),
		randPK(),
		randPK(),
	}
	alice   = pks[0]
	bob     = pks[1]
	charlie = pks[2]
	david   = pks[3]
	eve     = pks[4]
	faythe  = pks[5]
	grace   = pks[6]
	heidi   = pks[7]
)

func randPK() cryptotypes.PubKey {
	return ed25519.GenPrivKey().PubKey()
}

func pkAddr(pk cryptotypes.PubKey) string {
	return pkAccAddr(pk).String()
}

func pkAccAddr(pk cryptotypes.PubKey) sdk.AccAddress {
	return sdk.AccAddress(pk.Address())
}

// Prevent strconv unused error
var _ = strconv.IntSize

type SequencerTestSuite struct {
	apptesting.KeeperTestHelper
	msgServer   types.MsgServer
	queryClient types.QueryClient
}

func (s *SequencerTestSuite) k() *keeper.Keeper {
	return &s.App.SequencerKeeper
}

func (s *SequencerTestSuite) raK() *rollappkeeper.Keeper {
	return s.App.RollappKeeper
}

func TestSequencerKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(SequencerTestSuite))
}

func (s *SequencerTestSuite) SetupTest() {
	app := apptesting.Setup(s.T(), false)
	ctx := app.GetBaseApp().NewContext(false, cometbftproto.Header{})

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.SequencerKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	s.App = app
	s.msgServer = keeper.NewMsgServerImpl(app.SequencerKeeper)
	s.Ctx = ctx
	s.queryClient = queryClient
}

func (s *SequencerTestSuite) seq(pk cryptotypes.PubKey) types.Sequencer {
	return s.k().GetSequencer(s.Ctx, pkAddr(pk))
}

func (s *SequencerTestSuite) moduleBalance() sdk.Coin {
	acc := s.App.AccountKeeper.GetModuleAccount(s.Ctx, types.ModuleName)
	cs := s.App.BankKeeper.GetAllBalances(s.Ctx, acc.GetAddress())
	if cs.Len() == 0 {
		// coins will be zerod
		ret := bond
		ret.Amount = sdk.ZeroInt()
		return ret
	}
	return cs[0]
}

func (s *SequencerTestSuite) createRollapp() rollapptypes.Rollapp {
	return s.createRollappWithInitialSeqConstraint("*")
}

// init seq is an addr or empty or *
func (s *SequencerTestSuite) createRollappWithInitialSeqConstraint(initSeq string) rollapptypes.Rollapp {
	rollapp := rollapptypes.Rollapp{
		RollappId: urand.RollappID(),
		Owner:     sample.AccAddress(),
		GenesisInfo: rollapptypes.GenesisInfo{
			Bech32Prefix:    "rol",
			GenesisChecksum: "checksum",
			NativeDenom:     rollapptypes.DenomMetadata{Display: "DEN", Base: "aden", Exponent: 18},
			InitialSupply:   sdk.NewInt(1000),
		},
		InitialSequencer: initSeq,
	}
	s.raK().SetRollapp(s.Ctx, rollapp)
	return rollapp
}

func createSequencerMsg(rollapp string, pk cryptotypes.PubKey) types.MsgCreateSequencer {
	pkAny, err := codectypes.NewAnyWithValue(pk)
	if err != nil {
		panic(err)
	}

	return types.MsgCreateSequencer{
		Creator:      pkAddr(pk),
		DymintPubKey: pkAny,
		// Bond not included
		RollappId: rollapp,
		Metadata: types.SequencerMetadata{
			Rpcs:    []string{"https://rpc.wpd.evm.rollapp.noisnemyd.xyz:443"},
			EvmRpcs: []string{"https://rpc.evm.rollapp.noisnemyd.xyz:443"},
		},
	}
}

func (s *SequencerTestSuite) fundSequencer(pk cryptotypes.PubKey, amt sdk.Coin) {
	err := bankutil.FundAccount(s.App.BankKeeper, s.Ctx, pkAccAddr(pk), sdk.NewCoins(amt))
	s.Require().NoError(err)
}

func (s *SequencerTestSuite) createSequencerWithBond(ctx sdk.Context, rollapp string, pk cryptotypes.PubKey, bond sdk.Coin) types.Sequencer {
	s.fundSequencer(pk, bond)
	msg := createSequencerMsg(rollapp, pk)
	msg.Bond = bond
	_, err := s.msgServer.CreateSequencer(ctx, &msg)
	s.Require().NoError(err)
	return s.k().GetSequencer(ctx, pkAddr(pk))
}

func (s *SequencerTestSuite) equalSequencers(s1 *types.Sequencer, s2 *types.Sequencer) {
	eq := equalSequencers(s1, s2)
	s.Require().True(eq, "expected: %+v\nfound: %+v", *s1, *s2)
}

func equalSequencers(s1, s2 *types.Sequencer) bool {
	if s1.Address != s2.Address {
		return false
	}

	s1Pubkey := s1.DymintPubKey
	s2Pubkey := s2.DymintPubKey
	if !s1Pubkey.Equal(s2Pubkey) {
		return false
	}
	if s1.RollappId != s2.RollappId {
		return false
	}
	if !reflect.DeepEqual(s1.Metadata, s2.Metadata) {
		return false
	}

	if s1.Status != s2.Status {
		return false
	}

	if s1.OptedIn != s2.OptedIn {
		return false
	}

	if !s1.Tokens.IsEqual(s2.Tokens) {
		return false
	}

	if !s1.NoticePeriodTime.Equal(s2.NoticePeriodTime) {
		return false
	}

	return true
}

func createNSequencers(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Sequencer {
	items := make([]types.Sequencer, n)
	for i := range items {
		seq := types.Sequencer{
			Address: strconv.Itoa(i),
			Status:  types.Bonded,
		}
		items[i] = seq

		keeper.SetSequencer(ctx, items[i])
	}
	return items
}
