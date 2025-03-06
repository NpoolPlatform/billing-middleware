package addon

import (
	"context"

	addon1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/addon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/addon"
)

func (s *Server) CreateAddon(ctx context.Context, in *npool.CreateAddonRequest) (*npool.CreateAddonResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateAddon",
			"In", in,
		)
		return &npool.CreateAddonResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithEntID(req.EntID, false),
		addon1.WithAppID(req.AppID, true),
		addon1.WithPrice(req.Price, true),
		addon1.WithCredit(req.Credit, true),
		addon1.WithSortOrder(req.SortOrder, true),
		addon1.WithEnabled(req.Enabled, false),
		addon1.WithDescription(req.Description, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAddon",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateAddon(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateAddon",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAddonResponse{}, nil
}
