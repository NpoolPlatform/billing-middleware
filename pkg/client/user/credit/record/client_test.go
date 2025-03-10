package record

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
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/credit/record"
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

var ret = &npool.Record{
	EntID:            uuid.NewString(),
	AppID:            uuid.NewString(),
	UserID:           uuid.NewString(),
	OperationType:    types.OperationType_Exchange,
	OperationTypeStr: types.OperationType_Exchange.String(),
	CreditsChange:    int32(10),
	Extra:            uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createRecord(t *testing.T) {
	err := CreateRecord(context.Background(), &npool.RecordReq{
		EntID:         &ret.EntID,
		AppID:         &ret.AppID,
		UserID:        &ret.UserID,
		OperationType: &ret.OperationType,
		CreditsChange: &ret.CreditsChange,
		Extra:         &ret.Extra,
	})
	if assert.Nil(t, err) {
		info, err := GetRecord(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, ret)
		}
	}
}

func getRecord(t *testing.T) {
	info, err := GetRecord(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getRecords(t *testing.T) {
	conds := &npool.Conds{
		ID:            &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		OperationType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.OperationType)},
		IDs:           &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	}
	infos, err := GetRecords(context.Background(), conds, 0, 2)
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], ret)
	}
}

func deleteRecord(t *testing.T) {
	err := DeleteRecord(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetRecord(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestRecord(t *testing.T) {
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

	t.Run("createRecord", createRecord)
	t.Run("getRecord", getRecord)
	t.Run("getRecords", getRecords)
	t.Run("deleteRecord", deleteRecord)
}
