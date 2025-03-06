//nolint:dupl
package record

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	record1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/user/credit/record"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/credit/record"
)

func (s *Server) ExistRecordConds(ctx context.Context, in *npool.ExistRecordCondsRequest) (*npool.ExistRecordCondsResponse, error) {
	handler, err := record1.NewHandler(
		ctx,
		record1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRecordConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistRecordCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistRecordConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRecordConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistRecordCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistRecordCondsResponse{
		Info: exist,
	}, nil
}
