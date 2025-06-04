package subscription

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	subscription1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/user/subscription"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription"
)

func (s *Server) GetSubscription(ctx context.Context, in *npool.GetSubscriptionRequest) (*npool.GetSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetSubscription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriptionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSubscriptions(ctx context.Context, in *npool.GetSubscriptionsRequest) (*npool.GetSubscriptionsResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithConds(in.GetConds()),
		subscription1.WithOffset(in.GetOffset()),
		subscription1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.GetSubscriptions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriptionsResponse{
		Infos: infos,
	}, nil
}
