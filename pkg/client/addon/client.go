package addon

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	servicename "github.com/NpoolPlatform/billing-middleware/pkg/servicename"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/addon"
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

func CreateAddon(ctx context.Context, req *npool.AddonReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateAddon(_ctx, &npool.CreateAddonRequest{
			Info: req,
		})
	})
	return err
}

func GetAddon(ctx context.Context, id string) (*npool.Addon, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAddon(ctx, &npool.GetAddonRequest{
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
	return info.(*npool.Addon), nil
}

func GetAddons(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Addon, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAddons(ctx, &npool.GetAddonsRequest{
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
	return _infos.([]*npool.Addon), nil
}

func CountAddons(ctx context.Context, conds *npool.Conds) (total uint32, err error) {
	_, err = withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.CountAddons(ctx, &npool.CountAddonsRequest{
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

func ExistAddonConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistAddonConds(ctx, &npool.ExistAddonCondsRequest{
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

func GetAddonOnly(ctx context.Context, conds *npool.Conds) (*npool.Addon, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAddons(ctx, &npool.GetAddonsRequest{
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
	if len(infos.([]*npool.Addon)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Addon)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.Addon)[0], nil
}

func UpdateAddon(ctx context.Context, req *npool.AddonReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateAddon(_ctx, &npool.UpdateAddonRequest{
			Info: req,
		})
	})
	return err
}

func DeleteAddon(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteAddon(_ctx, &npool.DeleteAddonRequest{
			Info: &npool.AddonReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
