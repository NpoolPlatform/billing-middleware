package subscription

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/billing-middleware/pkg/testinit"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/billing/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &npool.Subscription{
	EntID:              uuid.NewString(),
	AppID:              uuid.NewString(),
	UserID:             uuid.NewString(),
	PackageID:          uuid.NewString(),
	StartAt:            uint32(10),
	EndAt:              uint32(20),
	OrderID:            uuid.NewString(),
	UsageState:         types.UsageState_Usful,
	UsageStateStr:      types.UsageState_Usful.String(),
	SubscriptionCredit: uint32(10),
	AddonCredit:        uint32(10),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createSubscription(t *testing.T) {
	err := CreateSubscription(context.Background(), &npool.SubscriptionReq{
		EntID:              &ret.EntID,
		AppID:              &ret.AppID,
		UserID:             &ret.UserID,
		PackageID:          &ret.PackageID,
		StartAt:            &ret.StartAt,
		EndAt:              &ret.EndAt,
		OrderID:            &ret.OrderID,
		UsageState:         &ret.UsageState,
		SubscriptionCredit: &ret.SubscriptionCredit,
		AddonCredit:        &ret.AddonCredit,
	})
	if assert.Nil(t, err) {
		info, err := GetSubscription(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, ret)
		}
	}
}

func updateSubscription(t *testing.T) {
	ret.StartAt = uint32(20)
	ret.EndAt = uint32(30)
	ret.UsageState = types.UsageState_Expire
	ret.UsageStateStr = types.UsageState_Expire.String()
	ret.SubscriptionCredit = uint32(20)
	ret.AddonCredit = uint32(30)
	err := UpdateSubscription(context.Background(), &npool.SubscriptionReq{
		ID:                 &ret.ID,
		EntID:              &ret.EntID,
		StartAt:            &ret.StartAt,
		EndAt:              &ret.EndAt,
		UsageState:         &ret.UsageState,
		SubscriptionCredit: &ret.SubscriptionCredit,
		AddonCredit:        &ret.AddonCredit,
	})
	if assert.Nil(t, err) {
		info, err := GetSubscription(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, ret)
		}
	}
}

func getSubscription(t *testing.T) {
	info, err := GetSubscription(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getSubscriptions(t *testing.T) {
	conds := &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		PackageID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.PackageID},
		OrderID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		StartAt:    &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.StartAt},
		EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.EndAt},
		UsageState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.UsageState)},
		IDs:        &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	}
	infos, err := GetSubscriptions(context.Background(), conds, 0, 2)
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], ret)
	}
}

func deleteSubscription(t *testing.T) {
	err := DeleteSubscription(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetSubscription(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestSubscription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createSubscription", createSubscription)
	t.Run("updateSubscription", updateSubscription)
	t.Run("getSubscription", getSubscription)
	t.Run("getSubscriptions", getSubscriptions)
	t.Run("deleteSubscription", deleteSubscription)
}
