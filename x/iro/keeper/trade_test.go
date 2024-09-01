package keeper_test

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/dymension/v3/testutil/sample"
	"github.com/dymensionxyz/dymension/v3/x/iro/types"
)

// FIXME: test trade after settled
// FIXME: test taker fee

func (s *KeeperTestSuite) TestBuy() {
	rollappId := s.CreateDefaultRollapp()
	k := s.App.IROKeeper

	startTime := time.Now().Add(time.Hour)
	maxAmt := sdk.NewInt(1_000_000_000)

	rollapp, _ := s.App.RollappKeeper.GetRollapp(s.Ctx, rollappId)
	planId, err := k.CreatePlan(s.Ctx, sdk.NewInt(1_000_000), startTime, startTime.Add(time.Hour), rollapp)
	s.Require().NoError(err)

	// buyer bech32
	buyer := sample.Acc()
	buyersFunds := sdk.NewCoins(sdk.NewCoin("adym", sdk.NewInt(100_000)))
	s.FundAcc(buyer, buyersFunds)

	// buy before plan start - should fail
	s.Ctx = s.Ctx.WithBlockTime(startTime.Add(-time.Minute))
	err = k.Buy(s.Ctx, planId, buyer.String(), sdk.NewInt(1_000), maxAmt)
	s.Require().Error(err)

	// plan start
	s.Ctx = s.Ctx.WithBlockTime(startTime.Add(time.Minute))

	// buy more than total allocation - should fail
	err = k.Buy(s.Ctx, planId, buyer.String(), sdk.NewInt(1_000_001), maxAmt)
	s.Require().Error(err)

	// cost is higher than maxCost specified - should fail
	err = k.Buy(s.Ctx, planId, buyer.String(), sdk.NewInt(1_000), sdk.NewInt(10))
	s.Require().Error(err)

	// buy more than user's balance - should fail
	err = k.Buy(s.Ctx, planId, buyer.String(), sdk.NewInt(900_000), maxAmt)
	s.Require().Error(err)

	// successful buy
	plan, _ := k.GetPlan(s.Ctx, planId)
	s.Assert().Equal(sdk.NewInt(0), plan.SoldAmt)

	amountTokensToBuy := sdk.NewInt(100)
	curve := types.NewBondingCurve(math.OneInt(), math.ZeroInt(), math.OneInt())
	expectedCost := curve.Cost(plan.SoldAmt, plan.SoldAmt.Add(amountTokensToBuy))
	err = k.Buy(s.Ctx, planId, buyer.String(), amountTokensToBuy, maxAmt)
	s.Require().NoError(err)
	plan, _ = k.GetPlan(s.Ctx, planId)
	s.Assert().True(plan.SoldAmt.Equal(amountTokensToBuy))

	// buy again, check cost changed
	expectedCost2 := curve.Cost(plan.SoldAmt, plan.SoldAmt.Add(amountTokensToBuy))
	err = k.Buy(s.Ctx, planId, buyer.String(), amountTokensToBuy, maxAmt)
	s.Require().NoError(err)
	s.Assert().True(expectedCost2.GT(expectedCost))

	// buy more than left - should fail
	err = k.Buy(s.Ctx, planId, buyer.String(), sdk.NewInt(999_999), maxAmt)
	s.Require().Error(err)

	// assert balance
	balances := s.App.BankKeeper.GetAllBalances(s.Ctx, buyer)
	s.Require().Equal(buyersFunds.AmountOf("adym").Sub(expectedCost).Sub(expectedCost2), balances.AmountOf("adym"))

	expectedBaseDenom := fmt.Sprintf("FUT_%s", rollappId)
	s.Require().Equal(amountTokensToBuy.MulRaw(2), balances.AmountOf(expectedBaseDenom))
}