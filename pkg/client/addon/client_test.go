package addon

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
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/addon"
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

var ret = &npool.Addon{
	EntID:       uuid.NewString(),
	AppID:       uuid.NewString(),
	Price:       decimal.NewFromInt(22).String(),
	Credit:      uint32(16),
	SortOrder:   uint32(1),
	Enabled:     true,
	Description: uuid.NewString(),
}

//nolint:unparam
func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createAddon(t *testing.T) {
	err := CreateAddon(context.Background(), &npool.AddonReq{
		EntID:       &ret.EntID,
		AppID:       &ret.AppID,
		Price:       &ret.Price,
		Credit:      &ret.Credit,
		SortOrder:   &ret.SortOrder,
		Enabled:     &ret.Enabled,
		Description: &ret.Description,
	})
	if assert.Nil(t, err) {
		info, err := GetAddon(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, ret)
		}
	}
}

func updateAddon(t *testing.T) {
	ret.Price = decimal.NewFromInt(15).String()
	ret.Credit = uint32(25)
	ret.SortOrder = uint32(2)
	ret.Enabled = false
	ret.Description = uuid.NewString()
	err := UpdateAddon(context.Background(), &npool.AddonReq{
		ID:          &ret.ID,
		EntID:       &ret.EntID,
		Price:       &ret.Price,
		Credit:      &ret.Credit,
		SortOrder:   &ret.SortOrder,
		Enabled:     &ret.Enabled,
		Description: &ret.Description,
	})
	if assert.Nil(t, err) {
		info, err := GetAddon(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, ret)
		}
	}
}

func getAddon(t *testing.T) {
	info, err := GetAddon(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getAddons(t *testing.T) {
	conds := &npool.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		SortOrder: &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.SortOrder},
		Enabled:   &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Enabled},
		IDs:       &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	}
	infos, err := GetAddons(context.Background(), conds, 0, 2)
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], ret)
	}
}

func deleteAddon(t *testing.T) {
	err := DeleteAddon(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetAddon(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestAddon(t *testing.T) {
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

	t.Run("createAddon", createAddon)
	t.Run("updateAddon", updateAddon)
	t.Run("getAddon", getAddon)
	t.Run("getAddons", getAddons)
	t.Run("deleteAddon", deleteAddon)
}
