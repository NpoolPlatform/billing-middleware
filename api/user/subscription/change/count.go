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

func (s *Server) CountSubscriptionChanges(
	ctx context.Context,
	in *npool.CountSubscriptionChangesRequest,
) (
	*npool.CountSubscriptionChangesResponse,
	error,
) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountSubscriptionChanges",
			"In", in,
			"Error", err,
		)
		return &npool.CountSubscriptionChangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	total, err := handler.CountSubscriptionChanges(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountSubscriptionChanges",
			"In", in,
			"Error", err,
		)
		return &npool.CountSubscriptionChangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CountSubscriptionChangesResponse{
		Total: total,
	}, nil
}
