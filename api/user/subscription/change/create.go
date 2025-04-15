package change

import (
	"context"

	subscription1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/user/subscription/change"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription/change"
)

func (s *Server) CreateSubscriptionChange(ctx context.Context, in *npool.CreateSubscriptionChangeRequest) (*npool.CreateSubscriptionChangeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateSubscriptionChange",
			"In", in,
		)
		return &npool.CreateSubscriptionChangeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithEntID(req.EntID, false),
		subscription1.WithAppID(req.AppID, true),
		subscription1.WithUserID(req.UserID, true),
		subscription1.WithUserSubscriptionID(req.UserSubscriptionID, true),
		subscription1.WithOldPackageID(req.OldPackageID, true),
		subscription1.WithNewPackageID(req.NewPackageID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSubscriptionChange",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSubscriptionChangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateSubscriptionChange(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateSubscriptionChange",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSubscriptionChangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSubscriptionChangeResponse{}, nil
}
