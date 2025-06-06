package addon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	addon1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/addon"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/addon"
)

func (s *Server) GetAddon(ctx context.Context, in *npool.GetAddonRequest) (*npool.GetAddonResponse, error) {
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAddon",
			"In", in,
			"Error", err,
		)
		return &npool.GetAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetAddon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAddon",
			"In", in,
			"Error", err,
		)
		return &npool.GetAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAddonResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAddons(ctx context.Context, in *npool.GetAddonsRequest) (*npool.GetAddonsResponse, error) {
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithConds(in.GetConds()),
		addon1.WithOffset(in.GetOffset()),
		addon1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAddons",
			"In", in,
			"Error", err,
		)
		return &npool.GetAddonsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.GetAddons(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAddons",
			"In", in,
			"Error", err,
		)
		return &npool.GetAddonsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAddonsResponse{
		Infos: infos,
	}, nil
}
