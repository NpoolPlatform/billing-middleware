package exchange

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	servicename "github.com/NpoolPlatform/billing-middleware/pkg/servicename"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
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

func CreateExchange(ctx context.Context, req *npool.ExchangeReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateExchange(_ctx, &npool.CreateExchangeRequest{
			Info: req,
		})
	})
	return err
}

func GetExchange(ctx context.Context, id string) (*npool.Exchange, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetExchange(ctx, &npool.GetExchangeRequest{
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
	return info.(*npool.Exchange), nil
}

func GetExchanges(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Exchange, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetExchanges(ctx, &npool.GetExchangesRequest{
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
	return _infos.([]*npool.Exchange), nil
}

func GetExchangesCount(ctx context.Context, conds *npool.Conds) (total uint32, err error) {
	_, err = withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetExchangesCount(ctx, &npool.GetExchangesCountRequest{
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

func ExistExchangeConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistExchangeConds(ctx, &npool.ExistExchangeCondsRequest{
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

func GetExchangeOnly(ctx context.Context, conds *npool.Conds) (*npool.Exchange, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetExchanges(ctx, &npool.GetExchangesRequest{
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
	if len(infos.([]*npool.Exchange)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Exchange)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.Exchange)[0], nil
}

func UpdateExchange(ctx context.Context, req *npool.ExchangeReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateExchange(_ctx, &npool.UpdateExchangeRequest{
			Info: req,
		})
	})
	return err
}

func DeleteExchange(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteExchange(_ctx, &npool.DeleteExchangeRequest{
			Info: &npool.ExchangeReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
