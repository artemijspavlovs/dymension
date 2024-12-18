package keeper_test

import (
	"cmp"
	"slices"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
	"github.com/dymensionxyz/sdk-utils/utils/urand"

	"github.com/dymensionxyz/dymension/v3/testutil/sample"
	"github.com/dymensionxyz/dymension/v3/x/rollapp/types"
)

func (s *RollappTestSuite) TestAddApp() {
	s.SetupTest()
	s.createRollappWithApp()
}

func (s *RollappTestSuite) createRollappWithApp() types.RollappSummary {
	s.SetupTest()
	creator := alice
	res := s.createRollappWithCreatorAndVerify(nil, creator, true)
	req := &types.MsgAddApp{
		Creator:     creator,
		Name:        "app1",
		RollappId:   res.RollappId,
		Description: "My first app",
		Image:       "http://example.com/image1",
		Url:         "http://example.com/app1",
		Order:       1,
	}
	_, err := s.msgServer.AddApp(s.Ctx, req)
	s.Require().NoError(err)

	// query the specific rollapp
	goCtx := sdk.WrapSDKContext(s.Ctx)
	queryResponse, err := s.queryClient.Rollapp(goCtx, &types.QueryGetRollappRequest{
		RollappId: req.GetRollappId(),
	})
	s.Require().Nil(err)

	appExpect := types.App{
		Id:          1,
		Name:        req.GetName(),
		RollappId:   req.GetRollappId(),
		Description: req.GetDescription(),
		ImageUrl:    req.GetImage(),
		Url:         req.GetUrl(),
		Order:       req.GetOrder(),
	}

	app, ok := s.k().GetApp(s.Ctx, 1, res.RollappId)
	s.Require().True(ok)
	s.Require().EqualValues(&appExpect, &app)
	s.Require().Len(queryResponse.Apps, 1)
	s.Require().EqualValues(&appExpect, queryResponse.Apps[0])

	return res
}

func (s *RollappTestSuite) Test_msgServer_AddApp() {
	rollappID := urand.RollappID()

	tests := []struct {
		name     string
		msgs     []*types.MsgAddApp
		malleate func()
		wantErr  error
	}{
		{
			name: "success: add 1 app",
			msgs: []*types.MsgAddApp{
				{
					Creator:     alice,
					Name:        "app1",
					RollappId:   rollappID,
					Description: "My first app",
					Image:       "http://example.com/image1",
					Url:         "http://example.com/app1",
					Order:       1,
				},
			},
		}, {
			name: "success: add 1 app no order",
			msgs: []*types.MsgAddApp{
				{
					Creator:     alice,
					Name:        "app1",
					RollappId:   rollappID,
					Description: "My first app",
					Image:       "http://example.com/image1",
					Url:         "http://example.com/app1",
				},
			},
		}, {
			name: "success: add multiple apps",
			msgs: []*types.MsgAddApp{
				{
					Creator:     alice,
					Name:        "app1",
					RollappId:   rollappID,
					Description: "My first app",
					Image:       "http://example.com/image1",
					Url:         "http://example.com/app1",
					Order:       3,
				}, {
					Creator:     alice,
					Name:        "app2",
					RollappId:   rollappID,
					Description: "My second app",
					Image:       "http://example.com/image2",
					Url:         "http://example.com/app2",
					Order:       1,
				}, {
					Creator:     alice,
					Name:        "app3",
					RollappId:   rollappID,
					Description: "My third app",
					Image:       "http://example.com/image3",
					Url:         "http://example.com/app3",
					Order:       4,
				},
			},
		}, {
			name: "fail: add app with different creator",
			msgs: []*types.MsgAddApp{
				{
					Creator:     bob,
					Name:        "app1",
					RollappId:   rollappID,
					Description: "My first app",
					Image:       "http://example.com/image1",
					Url:         "http://example.com/app1",
					Order:       1,
				},
			},
			wantErr: gerrc.ErrPermissionDenied,
		}, {
			name: "fail: add app with different rollapp",
			msgs: []*types.MsgAddApp{
				{
					Creator:     alice,
					Name:        "app1",
					RollappId:   urand.RollappID(),
					Description: "My first app",
					Image:       "http://example.com/image1",
					Url:         "http://example.com/app1",
					Order:       1,
				},
			},
			wantErr: types.ErrNotFound,
		}, {
			name: "fail: add app with same name and rollappID",
			msgs: []*types.MsgAddApp{
				{
					Creator:     alice,
					Name:        "app1",
					RollappId:   rollappID,
					Description: "My first app",
					Image:       "http://example.com/image1",
					Url:         "http://example.com/app1",
					Order:       1,
				},
			},
			malleate: func() {
				s.k().SetApp(s.Ctx, types.App{
					Name:      "app1",
					RollappId: rollappID,
				})
			},
			wantErr: gerrc.ErrAlreadyExists,
		}, {
			name: "success: add app with same name and different rollappID",
			msgs: []*types.MsgAddApp{
				{
					Creator:     alice,
					Name:        "app1",
					RollappId:   rollappID,
					Description: "My first app",
					Image:       "http://example.com/image1",
					Url:         "http://example.com/app1",
					Order:       1,
				},
			},
			malleate: func() {
				rollapp2 := urand.RollappID()
				s.createRollappWithIDAndCreator(rollapp2, alice)
				s.k().SetApp(s.Ctx, types.App{
					Name:      "app1",
					RollappId: rollapp2,
				})
			},
		}, {
			name: "fail: add 1 app - not enough funds",
			msgs: []*types.MsgAddApp{
				{
					Creator:     alice,
					Name:        "app1",
					RollappId:   rollappID,
					Description: "My first app",
					Image:       "http://example.com/image1",
					Url:         "http://example.com/app1",
					Order:       1,
				},
			},
			malleate: func() {
				params := s.k().GetParams(s.Ctx)
				params.AppRegistrationFee = sdk.NewInt64Coin("arax", 1)
				s.k().SetParams(s.Ctx, params)
			},
			wantErr: types.ErrAppRegistrationFeePayment,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s.SetupTest()
			s.createRollappWithIDAndCreator(rollappID, alice)

			goCtx := sdk.WrapSDKContext(s.Ctx)

			if tt.malleate != nil {
				tt.malleate()
			}

			for _, msg := range tt.msgs {
				err := func() error {
					err := msg.ValidateBasic()
					if err != nil {
						return err
					}
					_, err = s.msgServer.AddApp(goCtx, msg)
					return err
				}()
				if tt.wantErr != nil {
					s.Require().ErrorContains(err, tt.wantErr.Error())
				} else {
					s.Require().NoError(err)
				}
			}

			if tt.wantErr != nil {
				return
			}

			rollapp, err := s.queryClient.Rollapp(goCtx, &types.QueryGetRollappRequest{
				RollappId: rollappID,
			})
			s.Require().NoError(err)
			s.Require().Len(rollapp.Apps, len(tt.msgs))

			// check if the apps are ordered correctly
			slices.SortFunc(tt.msgs, func(a, b *types.MsgAddApp) int {
				return cmp.Compare(a.Order, b.Order)
			})
			for i, app := range rollapp.Apps {
				s.Require().Equal(tt.msgs[i].Order, app.Order)
			}

			// check if the app ids are sequenced correctly
			slices.SortFunc(rollapp.Apps, func(a, b *types.App) int {
				return int(a.Id) - int(b.Id)
			})
			for i, app := range rollapp.Apps {
				s.Require().Equal(uint64(i+1), app.Id)
			}
		})
	}
}

func (s *RollappTestSuite) Test_msgServer_UpdateApp() {
	rollappID := urand.RollappID()

	tests := []struct {
		name     string
		msgs     []*types.MsgUpdateApp
		malleate func()
		wantErr  error
	}{
		{
			name: "success: update existing app",
			msgs: []*types.MsgUpdateApp{
				{
					Id:          1,
					Creator:     alice,
					Name:        "app1",
					RollappId:   rollappID,
					Description: "Updated description",
					Image:       "http://example.com/updated_image",
					Url:         "http://example.com/updated_app",
					Order:       2,
				},
			},
			malleate: func() {
				s.k().SetApp(s.Ctx, types.App{
					Id:        1,
					Name:      "app1",
					RollappId: rollappID,
					Order:     1,
				})
			},
		}, {
			name: "fail: update non-existent app",
			msgs: []*types.MsgUpdateApp{
				{
					Id:          1,
					Creator:     alice,
					Name:        "non_existent_app",
					RollappId:   rollappID,
					Description: "This app does not exist",
					Image:       "http://example.com/non_existent_image",
					Url:         "http://example.com/non_existent_app",
					Order:       1,
				},
			},
			wantErr: gerrc.ErrNotFound,
		}, {
			name: "fail: update app with ID 0",
			msgs: []*types.MsgUpdateApp{
				{
					Id:          0,
					Creator:     alice,
					Name:        "app1",
					RollappId:   rollappID,
					Description: "Updated description",
					Image:       "http://example.com/updated_image",
					Url:         "http://example.com/updated_app",
					Order:       2,
				},
			},
			wantErr: types.ErrInvalidAppID,
		}, {
			name: "fail: update app with different creator",
			msgs: []*types.MsgUpdateApp{
				{
					Id:          1,
					Creator:     bob,
					Name:        "app1",
					RollappId:   rollappID,
					Description: "Trying to update with a different creator",
					Image:       "http://example.com/different_creator_image",
					Url:         "http://example.com/different_creator_app",
					Order:       2,
				},
			},
			malleate: func() {
				s.k().SetApp(s.Ctx, types.App{
					Id:        1,
					Name:      "app1",
					RollappId: rollappID,
					Order:     1,
				})
			},
			wantErr: gerrc.ErrPermissionDenied,
		}, {
			name: "success: update multiple apps",
			msgs: []*types.MsgUpdateApp{
				{
					Id:          1,
					Creator:     alice,
					Name:        "app1",
					RollappId:   rollappID,
					Description: "Updated app1",
					Image:       "http://example.com/updated_image1",
					Url:         "http://example.com/updated_app1",
					Order:       3,
				}, {
					Id:          2,
					Creator:     alice,
					Name:        "app2",
					RollappId:   rollappID,
					Description: "Updated app2",
					Image:       "http://example.com/updated_image2",
					Url:         "http://example.com/updated_app2",
					Order:       1,
				},
			},
			malleate: func() {
				s.k().SetApp(s.Ctx, types.App{
					Id:        1,
					Name:      "app1",
					RollappId: rollappID,
					Order:     2,
				})
				s.k().SetApp(s.Ctx, types.App{
					Id:        2,
					Name:      "app2",
					RollappId: rollappID,
					Order:     1,
				})
			},
		}, {
			name: "fail: update app with different rollapp",
			msgs: []*types.MsgUpdateApp{
				{
					Id:          1,
					Creator:     alice,
					Name:        "app1",
					RollappId:   urand.RollappID(),
					Description: "Trying to update with a different rollapp",
					Image:       "http://example.com/different_rollapp_image",
					Url:         "http://example.com/different_rollapp_app",
					Order:       1,
				},
			},
			malleate: func() {
				otherRollappID := urand.RollappID()
				s.createRollappWithIDAndCreator(otherRollappID, alice)
				s.k().SetApp(s.Ctx, types.App{
					Id:        1,
					Name:      "app1",
					RollappId: otherRollappID,
					Order:     1,
				})
			},
			wantErr: types.ErrNotFound,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s.SetupTest()
			s.createRollappWithIDAndCreator(rollappID, alice)

			goCtx := sdk.WrapSDKContext(s.Ctx)

			if tt.malleate != nil {
				tt.malleate()
			}

			for _, msg := range tt.msgs {
				err := func() error {
					err := msg.ValidateBasic()
					if err != nil {
						return err
					}
					_, err = s.msgServer.UpdateApp(goCtx, msg)
					return err
				}()
				if tt.wantErr != nil {
					s.Require().ErrorContains(err, tt.wantErr.Error())
				} else {
					s.Require().NoError(err)
				}
			}

			if tt.wantErr != nil {
				return
			}

			rollapp, err := s.queryClient.Rollapp(goCtx, &types.QueryGetRollappRequest{
				RollappId: rollappID,
			})
			s.Require().NoError(err)
			s.Require().Len(rollapp.Apps, len(tt.msgs))

			slices.SortFunc(tt.msgs, func(a, b *types.MsgUpdateApp) int {
				return cmp.Compare(a.Order, b.Order)
			})

			for i, app := range rollapp.Apps {
				s.Assert().Equal(tt.msgs[i].Order, app.Order)
				s.Assert().Equal(tt.msgs[i].Description, app.Description)
				s.Assert().Equal(tt.msgs[i].Image, app.ImageUrl)
				s.Assert().Equal(tt.msgs[i].Url, app.Url)
			}
		})
	}
}

func (s *RollappTestSuite) Test_msgServer_RemoveApp() {
	rollappID := urand.RollappID()

	tests := []struct {
		name     string
		msgs     []*types.MsgRemoveApp
		malleate func()
		wantErr  error
	}{
		{
			name: "success: remove existing app",
			msgs: []*types.MsgRemoveApp{
				{
					Creator:   alice,
					Id:        1,
					RollappId: rollappID,
				},
			},
			malleate: func() {
				s.k().SetApp(s.Ctx, types.App{
					Id:        1,
					Name:      "app1",
					RollappId: rollappID,
				})
			},
		}, {
			name: "fail: remove non-existent app",
			msgs: []*types.MsgRemoveApp{
				{
					Creator:   alice,
					Id:        144,
					RollappId: rollappID,
				},
			},
			wantErr: gerrc.ErrNotFound,
		}, {
			name: "fail: remove app with different creator",
			msgs: []*types.MsgRemoveApp{
				{
					Creator:   bob,
					Id:        1,
					RollappId: rollappID,
				},
			},
			malleate: func() {
				s.k().SetApp(s.Ctx, types.App{
					Id:        1,
					Name:      "app1",
					RollappId: rollappID,
				})
			},
			wantErr: gerrc.ErrPermissionDenied,
		}, {
			name: "fail: remove app with different rollapp",
			msgs: []*types.MsgRemoveApp{
				{
					Creator:   alice,
					Id:        1,
					RollappId: urand.RollappID(),
				},
			},
			malleate: func() {
				s.k().SetApp(s.Ctx, types.App{
					Id:        1,
					Name:      "app1",
					RollappId: rollappID,
				})
			},
			wantErr: types.ErrNotFound,
		}, {
			name: "success: remove multiple apps",
			msgs: []*types.MsgRemoveApp{
				{
					Creator:   alice,
					Id:        1,
					RollappId: rollappID,
				}, {
					Creator:   alice,
					Id:        2,
					RollappId: rollappID,
				},
			},
			malleate: func() {
				s.k().SetApp(s.Ctx, types.App{
					Id:        1,
					Name:      "app1",
					RollappId: rollappID,
				})
				s.k().SetApp(s.Ctx, types.App{
					Id:        2,
					Name:      "app2",
					RollappId: rollappID,
				})
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s.SetupTest()
			s.createRollappWithIDAndCreator(rollappID, alice)

			goCtx := sdk.WrapSDKContext(s.Ctx)

			if tt.malleate != nil {
				tt.malleate()
			}

			rollapp, err := s.queryClient.Rollapp(goCtx, &types.QueryGetRollappRequest{
				RollappId: rollappID,
			})
			s.Require().NoError(err)
			createdAppsCount := len(rollapp.Apps)

			for _, msg := range tt.msgs {
				err := func() error {
					err := msg.ValidateBasic()
					if err != nil {
						return err
					}
					_, err = s.msgServer.RemoveApp(goCtx, msg)
					return err
				}()
				if tt.wantErr != nil {
					s.Require().ErrorContains(err, tt.wantErr.Error())
				} else {
					s.Require().NoError(err)
				}
			}

			if tt.wantErr != nil {
				return
			}

			rollapp, err = s.queryClient.Rollapp(goCtx, &types.QueryGetRollappRequest{
				RollappId: rollappID,
			})
			s.Require().NoError(err)

			expectAppsCount := createdAppsCount - len(tt.msgs)
			s.Require().Len(rollapp.Apps, expectAppsCount)

			for _, msg := range tt.msgs {
				_, found := s.k().GetApp(s.Ctx, msg.GetId(), msg.RollappId)
				s.Require().False(found)
			}
		})
	}
}

func (s *RollappTestSuite) createRollappWithIDAndCreator(rollappId string, creator string) {
	rollapp := types.MsgCreateRollapp{
		Creator:          creator,
		RollappId:        rollappId,
		InitialSequencer: sample.AccAddress(),
		VmType:           types.Rollapp_EVM,
		Metadata:         &mockRollappMetadata,
		GenesisInfo:      mockGenesisInfo,
	}
	s.FundForAliasRegistration(rollapp)
	s.k().SetRollapp(s.Ctx, rollapp.GetRollapp())
}
