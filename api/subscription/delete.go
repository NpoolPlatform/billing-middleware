package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/subscription"
)

func (s *Server) DeleteSubscription(ctx context.Context, in *npool.DeleteSubscriptionRequest) (*npool.DeleteSubscriptionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteSubscription",
			"In", in,
		)
		return &npool.DeleteSubscriptionResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(req.ID, false),
		subscription1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteSubscription(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteSubscriptionResponse{}, nil
}
