package addon

import (
	"context"

	addon1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/addon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/addon"
)

func (s *Server) DeleteAddon(ctx context.Context, in *npool.DeleteAddonRequest) (*npool.DeleteAddonResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteAddon",
			"In", in,
		)
		return &npool.DeleteAddonResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithID(req.ID, false),
		addon1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAddon",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteAddon(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteAddon",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteAddonResponse{}, nil
}
