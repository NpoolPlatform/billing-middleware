package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/subscription"
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
		subscription1.WithPackageName(req.PackageName, false),
		subscription1.WithPrice(req.Price, false),
		subscription1.WithDescription(req.Description, false),
		subscription1.WithSortOrder(req.SortOrder, false),
		subscription1.WithCredit(req.Credit, false),
		subscription1.WithResetType(req.ResetType, false),
		subscription1.WithQpsLimit(req.QpsLimit, false),
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
