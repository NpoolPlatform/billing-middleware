package exchange

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
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
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

var ret = &npool.Exchange{
	EntID:             uuid.NewString(),
	AppID:             uuid.NewString(),
	UsageType:         types.UsageType_TextToken,
	UsageTypeStr:      types.UsageType_TextToken.String(),
	Credit:            uint32(1),
	ExchangeThreshold: uint32(1),
	Path:              uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createExchange(t *testing.T) {
	err := CreateExchange(context.Background(), &npool.ExchangeReq{
		EntID:             &ret.EntID,
		AppID:             &ret.AppID,
		UsageType:         &ret.UsageType,
		Credit:            &ret.Credit,
		ExchangeThreshold: &ret.ExchangeThreshold,
		Path:              &ret.Path,
	})
	if assert.Nil(t, err) {
		info, err := GetExchange(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, ret)
		}
	}
}

func updateExchange(t *testing.T) {
	ret.Credit = uint32(20)
	ret.ExchangeThreshold = uint32(15)
	ret.Path = uuid.NewString()
	err := UpdateExchange(context.Background(), &npool.ExchangeReq{
		ID:                &ret.ID,
		EntID:             &ret.EntID,
		Credit:            &ret.Credit,
		ExchangeThreshold: &ret.ExchangeThreshold,
		Path:              &ret.Path,
	})
	if assert.Nil(t, err) {
		info, err := GetExchange(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, ret)
		}
	}
}

func getExchange(t *testing.T) {
	info, err := GetExchange(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getExchanges(t *testing.T) {
	conds := &npool.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UsageType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.UsageType)},
		IDs:       &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		Path:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.Path},
	}
	infos, err := GetExchanges(context.Background(), conds, 0, 2)
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], ret)
	}
}

func deleteExchange(t *testing.T) {
	err := DeleteExchange(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetExchange(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestExchange(t *testing.T) {
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

	t.Run("createExchange", createExchange)
	t.Run("updateExchange", updateExchange)
	t.Run("getExchange", getExchange)
	t.Run("getExchanges", getExchanges)
	t.Run("deleteExchange", deleteExchange)
}
