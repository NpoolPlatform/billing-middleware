package addon

import (
	"context"

	addon1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/addon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/addon"
)

func (s *Server) UpdateAddon(ctx context.Context, in *npool.UpdateAddonRequest) (*npool.UpdateAddonResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateAddon",
			"In", in,
		)
		return &npool.UpdateAddonResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithID(req.ID, false),
		addon1.WithEntID(req.EntID, false),
		addon1.WithPrice(req.Price, false),
		addon1.WithCredit(req.Credit, false),
		addon1.WithSortOrder(req.SortOrder, false),
		addon1.WithEnabled(req.Enabled, false),
		addon1.WithDescription(req.Description, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAddon",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateAddon(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateAddon",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateAddonResponse{}, nil
}
