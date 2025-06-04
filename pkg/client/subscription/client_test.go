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
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/subscription"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
	EntID:          uuid.NewString(),
	AppID:          uuid.NewString(),
	PackageName:    uuid.NewString(),
	UsdPrice:       decimal.NewFromInt(10).String(),
	Description:    uuid.NewString(),
	SortOrder:      uint32(1),
	PackageType:    types.PackageType_Normal,
	PackageTypeStr: types.PackageType_Normal.String(),
	Credit:         uint32(10),
	ResetType:      types.ResetType_Monthly,
	ResetTypeStr:   types.ResetType_Monthly.String(),
	QPSLimit:       uint32(1),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createSubscription(t *testing.T) {
	err := CreateSubscription(context.Background(), &npool.SubscriptionReq{
		EntID:       &ret.EntID,
		AppID:       &ret.AppID,
		PackageName: &ret.PackageName,
		UsdPrice:    &ret.UsdPrice,
		Description: &ret.Description,
		SortOrder:   &ret.SortOrder,
		PackageType: &ret.PackageType,
		Credit:      &ret.Credit,
		ResetType:   &ret.ResetType,
		QPSLimit:    &ret.QPSLimit,
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
	ret.PackageName = uuid.NewString()
	ret.UsdPrice = decimal.NewFromInt(10).String()
	ret.Description = uuid.NewString()
	ret.SortOrder = uint32(2)
	ret.Credit = uint32(10)
	ret.ResetType = types.ResetType_Quarterly
	ret.ResetTypeStr = types.ResetType_Quarterly.String()
	ret.QPSLimit = uint32(5)
	err := UpdateSubscription(context.Background(), &npool.SubscriptionReq{
		ID:          &ret.ID,
		EntID:       &ret.EntID,
		PackageName: &ret.PackageName,
		UsdPrice:    &ret.UsdPrice,
		Description: &ret.Description,
		SortOrder:   &ret.SortOrder,
		Credit:      &ret.Credit,
		ResetType:   &ret.ResetType,
		QPSLimit:    &ret.QPSLimit,
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
		ID:          &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		PackageName: &basetypes.StringVal{Op: cruder.EQ, Value: ret.PackageName},
		SortOrder:   &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.SortOrder},
		PackageType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.PackageType)},
		ResetType:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.ResetType)},
		IDs:         &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntIDs:      &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
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
