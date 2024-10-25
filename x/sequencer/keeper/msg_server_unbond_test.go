package keeper_test

import (
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/dymensionxyz/dymension/v3/x/sequencer/types"
)

func (s *SequencerTestSuite) TestUnbondingNonProposer() {
	rollappId, pk := s.createRollapp()
	proposerAddr := s.createSequencer(s.Ctx, rollappId, pk)

	bondedAddr := s.CreateDefaultSequencer(s.Ctx, rollappId)
	s.Require().NotEqual(proposerAddr, bondedAddr)

	proposer := s.k().GetProposer(s.Ctx, rollappId)
	s.Require().True(ok)
	s.Equal(proposerAddr, proposer.Address)

	/* ------------------------- unbond non proposer sequencer ------------------------ */
	bondedSeq, err := s.k().GetRealSequencer(s.Ctx, bondedAddr)
	s.Require().True(found)
	s.Equal(types.Bonded, bondedSeq.Status)

	unbondMsg := types.MsgUnbond{Creator: bondedAddr}
	_, err := s.msgServer.Unbond(s.Ctx, &unbondMsg)
	s.Require().NoError(err)

	// check sequencer operating status
	bondedSeq, err = s.k().GetRealSequencer(s.Ctx, bondedAddr)
	s.Require().True(found)
	s.Equal(types.Unbonding, bondedSeq.Status)

	s.k().UnbondAllMatureSequencers(s.Ctx, bondedSeq.UnbondTime.Add(10*time.Second))
	bondedSeq, err = s.k().GetRealSequencer(s.Ctx, bondedAddr)
	s.Require().True(found)
	s.Equal(types.Unbonded, bondedSeq.Status)

	// check proposer not changed
	proposer, ok = s.k().GetProposerLegacy(s.Ctx, rollappId)
	s.Require().True(ok)
	s.Equal(proposerAddr, proposer.Address)

	// try to unbond again. already unbonded, we expect error
	_, err = s.msgServer.Unbond(s.Ctx, &unbondMsg)
	s.Require().Error(err)
}

func (s *SequencerTestSuite) TestUnbondingProposer() {
	s.Ctx = s.Ctx.WithBlockHeight(10)

	rollappId, proposerAddr := s.CreateDefaultRollappAndProposer()
	_ = s.createSequencer(s.Ctx, rollappId, ed25519.GenPrivKey().PubKey())

	/* ----------------------------- unbond proposer ---------------------------- */
	unbondMsg := types.MsgUnbond{Creator: proposerAddr}
	_, err := s.msgServer.Unbond(s.Ctx, &unbondMsg)
	s.Require().NoError(err)

	// check proposer still bonded and notice period started
	p := s.k().GetProposer(s.Ctx, rollappId)
	s.Require().True(ok)
	s.Equal(proposerAddr, p.Address)
	s.Equal(s.Ctx.BlockHeight(), p.UnbondRequestHeight)

	// unbonding again, we expect error as sequencer is in notice period
	_, err = s.msgServer.Unbond(s.Ctx, &unbondMsg)
	s.Require().Error(err)

	// next proposer should not be set yet
	_, ok = s.k().GetNextProposer(s.Ctx, rollappId)
	s.Require().False(ok)

	// check notice period queue
	m, err := s.k().NoticeElapsedSequencers(s.Ctx, p.NoticePeriodTime.Add(-1*time.Second))
	s.Require().NoError(err)
	s.Require().Len(m, 0)
	m, err = s.k().NoticeElapsedSequencers(s.Ctx, p.NoticePeriodTime.Add(1*time.Second))
	s.Require().NoError(err)
	s.Require().Len(m, 1)
}
