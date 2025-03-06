package exchange

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/billing-middleware/pkg/testinit"
	types "github.com/NpoolPlatform/message/npool/basetypes/billing/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.Exchange{
	EntID:             uuid.NewString(),
	AppID:             uuid.NewString(),
	UsageType:         types.UsageType_TextToken,
	UsageTypeStr:      types.UsageType_TextToken.String(),
	Credit:            uint32(1),
	ExchangeThreshold: uint32(1),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createExchange(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUsageType(&ret.UsageType, true),
		WithCredit(&ret.Credit, true),
		WithExchangeThreshold(&ret.ExchangeThreshold, true),
	)
	assert.Nil(t, err)

	err = handler.CreateExchange(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetExchange(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateExchange(t *testing.T) {
	ret.Credit = uint32(10)
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithUsageType(&ret.UsageType, true),
		WithCredit(&ret.Credit, true),
		WithExchangeThreshold(&ret.ExchangeThreshold, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateExchange(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetExchange(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getExchange(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetExchange(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getExchanges(t *testing.T) {
	conds := &npool.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UsageType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.UsageType)},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, err := handler.GetExchanges(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteExchange(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteExchange(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetExchange(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestExchange(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createExchange", createExchange)
	t.Run("updateExchange", updateExchange)
	t.Run("getExchange", getExchange)
	t.Run("getExchanges", getExchanges)
	t.Run("deleteExchange", deleteExchange)
}
