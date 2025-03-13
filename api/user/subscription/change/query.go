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

func (s *Server) GetSubscriptionChange(ctx context.Context, in *npool.GetSubscriptionChangeRequest) (*npool.GetSubscriptionChangeResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionChange",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionChangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetSubscriptionChange(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionChange",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionChangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriptionChangeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSubscriptionChanges(ctx context.Context, in *npool.GetSubscriptionChangesRequest) (*npool.GetSubscriptionChangesResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithConds(in.GetConds()),
		subscription1.WithOffset(in.GetOffset()),
		subscription1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionChanges",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionChangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.GetSubscriptionChanges(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionChanges",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionChangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriptionChangesResponse{
		Infos: infos,
	}, nil
}

func (s *Server) GetSubscriptionChangesCount(
	ctx context.Context,
	in *npool.GetSubscriptionChangesCountRequest,
) (
	*npool.GetSubscriptionChangesCountResponse,
	error,
) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionChangesCount",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionChangesCountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	total, err := handler.GetSubscriptionChangesCount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionChangesCount",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionChangesCountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriptionChangesCountResponse{
		Total: total,
	}, nil
}
