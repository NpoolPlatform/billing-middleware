package change

import (
	"context"

	subscription1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/user/subscription/change"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription/change"
)

func (s *Server) DeleteSubscriptionChange(ctx context.Context, in *npool.DeleteSubscriptionChangeRequest) (*npool.DeleteSubscriptionChangeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteSubscriptionChange",
			"In", in,
		)
		return &npool.DeleteSubscriptionChangeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(req.ID, false),
		subscription1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSubscriptionChange",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSubscriptionChangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteSubscriptionChange(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteSubscriptionChange",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSubscriptionChangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteSubscriptionChangeResponse{}, nil
}
