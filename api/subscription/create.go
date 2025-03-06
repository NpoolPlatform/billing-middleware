package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/subscription"
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
		subscription1.WithPackageName(req.PackageName, true),
		subscription1.WithPrice(req.Price, true),
		subscription1.WithDescription(req.Description, false),
		subscription1.WithSortOrder(req.SortOrder, true),
		subscription1.WithPackageType(req.PackageType, true),
		subscription1.WithCredit(req.Credit, true),
		subscription1.WithResetType(req.ResetType, true),
		subscription1.WithQpsLimit(req.QpsLimit, true),
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
