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

func (s *Server) ExistAddonConds(ctx context.Context, in *npool.ExistAddonCondsRequest) (*npool.ExistAddonCondsResponse, error) {
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAddonConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAddonCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistAddonConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAddonConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAddonCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistAddonCondsResponse{
		Info: exist,
	}, nil
}
