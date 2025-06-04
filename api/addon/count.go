//nolint:dupl
package addon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	addon1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/addon"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/addon"
)

func (s *Server) CountAddons(ctx context.Context, in *npool.CountAddonsRequest) (*npool.CountAddonsResponse, error) {
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountAddons",
			"In", in,
			"Error", err,
		)
		return &npool.CountAddonsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	total, err := handler.CountAddons(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountAddons",
			"In", in,
			"Error", err,
		)
		return &npool.CountAddonsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CountAddonsResponse{
		Total: total,
	}, nil
}
