package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/user/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription"
)

func (s *Server) UpdateSubscription(ctx context.Context, in *npool.UpdateSubscriptionRequest) (*npool.UpdateSubscriptionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateSubscription",
			"In", in,
		)
		return &npool.UpdateSubscriptionResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(req.ID, false),
		subscription1.WithEntID(req.EntID, false),
		subscription1.WithStartAt(req.StartAt, false),
		subscription1.WithEndAt(req.EndAt, false),
		subscription1.WithPackageID(req.PackageID, false),
		subscription1.WithUsageState(req.UsageState, false),
		subscription1.WithSubscriptionCredit(req.SubscriptionCredit, false),
		subscription1.WithAddonCredit(req.AddonCredit, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateSubscription(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateSubscriptionResponse{}, nil
}
