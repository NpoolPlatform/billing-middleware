package record

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	servicename "github.com/NpoolPlatform/billing-middleware/pkg/servicename"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/credit/record"
	"google.golang.org/grpc"
)

func withClient(ctx context.Context, handler func(context.Context, npool.MiddlewareClient) (interface{}, error)) (interface{}, error) {
	return grpc2.WithGRPCConn(
		ctx,
		servicename.ServiceDomain,
		10*time.Second, //nolint
		func(_ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
			return handler(_ctx, npool.NewMiddlewareClient(conn))
		},
		grpc2.GRPCTAG,
	)
}

func CreateRecord(ctx context.Context, req *npool.RecordReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateRecord(_ctx, &npool.CreateRecordRequest{
			Info: req,
		})
	})
	return err
}

func GetRecord(ctx context.Context, id string) (*npool.Record, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetRecord(ctx, &npool.GetRecordRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Record), nil
}

func GetRecords(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Record, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetRecords(ctx, &npool.GetRecordsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return _infos.([]*npool.Record), nil
}

func GetRecordsCount(ctx context.Context, conds *npool.Conds) (total uint32, err error) {
	_, err = withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetRecordsCount(ctx, &npool.GetRecordsCountRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		total = resp.Total
		return resp.Total, nil
	})
	if err != nil {
		return 0, err
	}
	return total, nil
}

func ExistRecordConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistRecordConds(ctx, &npool.ExistRecordCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func GetRecordOnly(ctx context.Context, conds *npool.Conds) (*npool.Record, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetRecords(ctx, &npool.GetRecordsRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.Record)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Record)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.Record)[0], nil
}

func DeleteRecord(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteRecord(_ctx, &npool.DeleteRecordRequest{
			Info: &npool.RecordReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
