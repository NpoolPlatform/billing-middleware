//nolint:dupl
package subscription

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	subscription1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/user/subscription"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription"
)

func (s *Server) ExistSubscriptionConds(
	ctx context.Context,
	in *npool.ExistSubscriptionCondsRequest,
) (
	*npool.ExistSubscriptionCondsResponse,
	error,
) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSubscriptionConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSubscriptionCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistSubscriptionConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSubscriptionConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSubscriptionCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistSubscriptionCondsResponse{
		Info: exist,
	}, nil
}
