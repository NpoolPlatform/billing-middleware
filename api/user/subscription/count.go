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

func (s *Server) CountSubscriptions(
	ctx context.Context,
	in *npool.CountSubscriptionsRequest,
) (
	*npool.CountSubscriptionsResponse,
	error,
) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountSubscriptions",
			"In", in,
			"Error", err,
		)
		return &npool.CountSubscriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	total, err := handler.CountSubscriptions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountSubscriptions",
			"In", in,
			"Error", err,
		)
		return &npool.CountSubscriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CountSubscriptionsResponse{
		Total: total,
	}, nil
}
