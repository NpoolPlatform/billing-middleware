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

func (s *Server) CountRecords(ctx context.Context, in *npool.CountRecordsRequest) (*npool.CountRecordsResponse, error) {
	handler, err := record1.NewHandler(
		ctx,
		record1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountRecords",
			"In", in,
			"Error", err,
		)
		return &npool.CountRecordsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	total, err := handler.CountRecords(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountRecords",
			"In", in,
			"Error", err,
		)
		return &npool.CountRecordsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CountRecordsResponse{
		Total: total,
	}, nil
}
