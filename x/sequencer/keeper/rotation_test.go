package keeper_test

import (
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/dymension/v3/x/sequencer/types"
)

func (s *SequencerTestSuite) TestExpectedNextProposer() {
	type testCase struct {
		name                    string
		numSeqAddrs             int
		expectEmptyNextProposer bool
	}

	testCases := []testCase{
		{"No additional sequencers", 0, true},
		{"few", 4, false},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()

			rollappId, pk := s.CreateDefaultRollapp()
			_ = s.CreateSequencerWithBond(s.Ctx, rollappId, bond, pk) // proposer, with highest bond

			seqAddrs := make([]string, tc.numSeqAddrs)
			currBond := sdk.NewCoin(bond.Denom, bond.Amount.Quo(sdk.NewInt(10)))
			for i := 0; i < len(seqAddrs); i++ {
				currBond = currBond.AddAmount(bond.Amount)
				pubkey := ed25519.GenPrivKey().PubKey()
				seqAddrs[i] = s.CreateSequencerWithBond(s.Ctx, rollappId, currBond, pubkey)
			}
			next := s.App.SequencerKeeper.ExpectedNextProposer(s.Ctx, rollappId)
			if tc.expectEmptyNextProposer {
				s.Require().Empty(next.Address)
				return
			}

			expectedNextProposer := seqAddrs[len(seqAddrs)-1]
			s.Equal(expectedNextProposer, next.Address)
		})
	}
}

// TestStartRotation tests the StartRotation function which is called when a sequencer has finished its notice period
func (s *SequencerTestSuite) TestStartRotation() {
	rollappId, pk := s.CreateDefaultRollapp()
	addr1 := s.CreateSequencer(s.Ctx, rollappId, pk)

	_ = s.CreateDefaultSequencer(s.Ctx, rollappId)
	_ = s.CreateDefaultSequencer(s.Ctx, rollappId)

	/* ----------------------------- unbond proposer ---------------------------- */
	unbondMsg := types.MsgUnbond{Creator: addr1}
	_, err := s.msgServer.Unbond(s.Ctx, &unbondMsg)
	s.Require().NoError(err)

	// check proposer still bonded and notice period started
	p, ok := s.App.SequencerKeeper.GetProposerLegacy(s.Ctx, rollappId)
	s.Require().True(ok)
	s.Equal(addr1, p.Address)
	s.Equal(s.Ctx.BlockHeight(), p.UnbondRequestHeight)

	m := s.App.SequencerKeeper.GetMatureNoticePeriodSequencers(s.Ctx, p.NoticePeriodTime.Add(-10*time.Second))
	s.Require().Len(m, 0)
	m = s.App.SequencerKeeper.GetMatureNoticePeriodSequencers(s.Ctx, p.NoticePeriodTime.Add(10*time.Second))
	s.Require().Len(m, 1)
	s.App.SequencerKeeper.MatureSequencersWithNoticePeriod(s.Ctx, p.NoticePeriodTime.Add(10*time.Second))

	// validate nextProposer is set
	n, ok := s.App.SequencerKeeper.GetNextProposer(s.Ctx, rollappId)
	s.Require().True(ok)
	s.Require().NotEmpty(n.Address)

	// validate proposer not changed
	p, _ = s.App.SequencerKeeper.GetProposerLegacy(s.Ctx, rollappId)
	s.Equal(addr1, p.Address)
}

func (s *SequencerTestSuite) TestRotateProposer() {
	rollappId, pk := s.CreateDefaultRollapp()
	addr1 := s.CreateSequencer(s.Ctx, rollappId, pk)
	addr2 := s.CreateSequencer(s.Ctx, rollappId, ed25519.GenPrivKey().PubKey())

	/* ----------------------------- unbond proposer ---------------------------- */
	unbondMsg := types.MsgUnbond{Creator: addr1}
	res, err := s.msgServer.Unbond(s.Ctx, &unbondMsg)
	s.Require().NoError(err)

	// mature notice period
	s.App.SequencerKeeper.MatureSequencersWithNoticePeriod(s.Ctx, res.GetNoticePeriodCompletionTime().Add(10*time.Second))
	_, ok := s.App.SequencerKeeper.GetNextProposer(s.Ctx, rollappId)
	s.Require().True(ok)

	// simulate lastBlock received
	err = s.App.SequencerKeeper.completeRotationLeg(s.Ctx, rollappId)
	s.Require().NoError(err)

	// assert addr2 is now proposer
	p, ok := s.App.SequencerKeeper.GetProposerLegacy(s.Ctx, rollappId)
	s.Require().True(ok)
	s.Equal(addr2, p.Address)
	// assert addr1 is unbonding
	u, _ := s.App.SequencerKeeper.GetSequencer(s.Ctx, addr1)
	s.Equal(types.Unbonding, u.Status)
	// assert nextProposer is nil
	_, ok = s.App.SequencerKeeper.GetNextProposer(s.Ctx, rollappId)
	s.Require().False(ok)
}

func (s *SequencerTestSuite) TestRotateProposerNoNextProposer() {
	rollappId, pk := s.CreateDefaultRollapp()
	addr1 := s.CreateSequencer(s.Ctx, rollappId, pk)

	/* ----------------------------- unbond proposer ---------------------------- */
	unbondMsg := types.MsgUnbond{Creator: addr1}
	res, err := s.msgServer.Unbond(s.Ctx, &unbondMsg)
	s.Require().NoError(err)

	// mature notice period
	s.App.SequencerKeeper.MatureSequencersWithNoticePeriod(s.Ctx, res.GetNoticePeriodCompletionTime().Add(10*time.Second))
	// simulate lastBlock received
	err = s.App.SequencerKeeper.completeRotationLeg(s.Ctx, rollappId)
	s.Require().NoError(err)

	_, ok := s.App.SequencerKeeper.GetProposerLegacy(s.Ctx, rollappId)
	s.Require().False(ok)

	_, ok = s.App.SequencerKeeper.GetNextProposer(s.Ctx, rollappId)
	s.Require().False(ok)
}

// Both the proposer and nextProposer tries to unbond
func (s *SequencerTestSuite) TestStartRotationTwice() {
	s.Ctx = s.Ctx.WithBlockHeight(10)

	rollappId, pk := s.CreateDefaultRollapp()
	addr1 := s.CreateSequencer(s.Ctx, rollappId, pk)
	addr2 := s.CreateSequencer(s.Ctx, rollappId, ed25519.GenPrivKey().PubKey())

	// unbond proposer
	unbondMsg := types.MsgUnbond{Creator: addr1}
	_, err := s.msgServer.Unbond(s.Ctx, &unbondMsg)
	s.Require().NoError(err)

	p, ok := s.App.SequencerKeeper.GetProposerLegacy(s.Ctx, rollappId)
	s.Require().True(ok)
	s.Equal(addr1, p.Address)
	s.Equal(s.Ctx.BlockHeight(), p.UnbondRequestHeight)

	s.App.SequencerKeeper.MatureSequencersWithNoticePeriod(s.Ctx, p.NoticePeriodTime.Add(10*time.Second))
	s.Require().True(s.App.SequencerKeeper.isRotatingLeg(s.Ctx, rollappId))

	n, ok := s.App.SequencerKeeper.GetNextProposer(s.Ctx, rollappId)
	s.Require().True(ok)
	s.Equal(addr2, n.Address)

	// unbond nextProposer before rotation completes
	s.Ctx = s.Ctx.WithBlockHeight(20)
	unbondMsg = types.MsgUnbond{Creator: addr2}
	_, err = s.msgServer.Unbond(s.Ctx, &unbondMsg)
	s.Require().NoError(err)

	// check nextProposer is still the nextProposer and notice period started
	n, ok = s.App.SequencerKeeper.GetNextProposer(s.Ctx, rollappId)
	s.Require().True(ok)
	s.Equal(addr2, n.Address)
	s.Require().True(n.IsNoticePeriodInProgress())

	// rotation completes before notice period ends for addr2 (the nextProposer)
	err = s.App.SequencerKeeper.completeRotationLeg(s.Ctx, rollappId) // simulate lastBlock received
	s.Require().NoError(err)

	// validate addr2 is now proposer and still with notice period
	p, _ = s.App.SequencerKeeper.GetProposerLegacy(s.Ctx, rollappId)
	s.Equal(addr2, p.Address)
	s.Require().True(p.IsNoticePeriodInProgress())

	// validate nextProposer is unset after rotation completes
	n, ok = s.App.SequencerKeeper.GetNextProposer(s.Ctx, rollappId)
	s.Require().False(ok)

	// mature notice period for addr2
	s.App.SequencerKeeper.MatureSequencersWithNoticePeriod(s.Ctx, p.NoticePeriodTime.Add(10*time.Second))
	// validate nextProposer is set
	n, ok = s.App.SequencerKeeper.GetNextProposer(s.Ctx, rollappId)
	s.Require().True(ok)
	s.Require().Empty(n.Address)
}
