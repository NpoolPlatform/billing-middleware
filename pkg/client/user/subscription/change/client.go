package change

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	servicename "github.com/NpoolPlatform/billing-middleware/pkg/servicename"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription/change"
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

func CreateSubscriptionChange(ctx context.Context, req *npool.SubscriptionChangeReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateSubscriptionChange(_ctx, &npool.CreateSubscriptionChangeRequest{
			Info: req,
		})
	})
	return err
}

func GetSubscriptionChange(ctx context.Context, id string) (*npool.SubscriptionChange, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetSubscriptionChange(ctx, &npool.GetSubscriptionChangeRequest{
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
	return info.(*npool.SubscriptionChange), nil
}

func GetSubscriptionChanges(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.SubscriptionChange, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetSubscriptionChanges(ctx, &npool.GetSubscriptionChangesRequest{
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
	return _infos.([]*npool.SubscriptionChange), nil
}

func CountSubscriptionChanges(ctx context.Context, conds *npool.Conds) (total uint32, err error) {
	_, err = withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.CountSubscriptionChanges(ctx, &npool.CountSubscriptionChangesRequest{
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

func ExistSubscriptionChangeConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistSubscriptionChangeConds(ctx, &npool.ExistSubscriptionChangeCondsRequest{
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

func GetSubscriptionChangeOnly(ctx context.Context, conds *npool.Conds) (*npool.SubscriptionChange, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetSubscriptionChanges(ctx, &npool.GetSubscriptionChangesRequest{
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
	if len(infos.([]*npool.SubscriptionChange)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.SubscriptionChange)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.SubscriptionChange)[0], nil
}

func DeleteSubscriptionChange(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteSubscriptionChange(_ctx, &npool.DeleteSubscriptionChangeRequest{
			Info: &npool.SubscriptionChangeReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
