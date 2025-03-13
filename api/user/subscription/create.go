package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/user/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription"
)

func (s *Server) CreateSubscription(ctx context.Context, in *npool.CreateSubscriptionRequest) (*npool.CreateSubscriptionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateSubscription",
			"In", in,
		)
		return &npool.CreateSubscriptionResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithEntID(req.EntID, false),
		subscription1.WithAppID(req.AppID, true),
		subscription1.WithUserID(req.UserID, true),
		subscription1.WithPackageID(req.PackageID, true),
		subscription1.WithStartAt(req.StartAt, true),
		subscription1.WithEndAt(req.EndAt, true),
		subscription1.WithUsageState(req.UsageState, true),
		subscription1.WithSubscriptionCredit(req.SubscriptionCredit, true),
		subscription1.WithAddonCredit(req.AddonCredit, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateSubscription(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSubscriptionResponse{}, nil
}
