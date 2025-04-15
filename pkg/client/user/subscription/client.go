package subscription

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	servicename "github.com/NpoolPlatform/billing-middleware/pkg/servicename"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription"
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

func CreateSubscription(ctx context.Context, req *npool.SubscriptionReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateSubscription(_ctx, &npool.CreateSubscriptionRequest{
			Info: req,
		})
	})
	return err
}

func GetSubscription(ctx context.Context, id string) (*npool.Subscription, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetSubscription(ctx, &npool.GetSubscriptionRequest{
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
	return info.(*npool.Subscription), nil
}

func GetSubscriptions(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Subscription, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetSubscriptions(ctx, &npool.GetSubscriptionsRequest{
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
	return _infos.([]*npool.Subscription), nil
}

func GetSubscriptionsCount(ctx context.Context, conds *npool.Conds) (total uint32, err error) {
	_, err = withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetSubscriptionsCount(ctx, &npool.GetSubscriptionsCountRequest{
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

func ExistSubscriptionConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistSubscriptionConds(ctx, &npool.ExistSubscriptionCondsRequest{
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

func GetSubscriptionOnly(ctx context.Context, conds *npool.Conds) (*npool.Subscription, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetSubscriptions(ctx, &npool.GetSubscriptionsRequest{
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
	if len(infos.([]*npool.Subscription)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Subscription)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.Subscription)[0], nil
}

func UpdateSubscription(ctx context.Context, req *npool.SubscriptionReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateSubscription(_ctx, &npool.UpdateSubscriptionRequest{
			Info: req,
		})
	})
	return err
}

func DeleteSubscription(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteSubscription(_ctx, &npool.DeleteSubscriptionRequest{
			Info: &npool.SubscriptionReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
