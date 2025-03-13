//nolint:dupl
package change

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	subscription1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/user/subscription/change"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription/change"
)

func (s *Server) ExistSubscriptionChangeConds(
	ctx context.Context,
	in *npool.ExistSubscriptionChangeCondsRequest,
) (
	*npool.ExistSubscriptionChangeCondsResponse,
	error,
) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSubscriptionChangeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSubscriptionChangeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistSubscriptionChangeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSubscriptionChangeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSubscriptionChangeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistSubscriptionChangeCondsResponse{
		Info: exist,
	}, nil
}
