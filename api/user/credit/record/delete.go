package record

import (
	"context"

	record1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/user/credit/record"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/credit/record"
)

func (s *Server) DeleteRecord(ctx context.Context, in *npool.DeleteRecordRequest) (*npool.DeleteRecordResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteRecord",
			"In", in,
		)
		return &npool.DeleteRecordResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := record1.NewHandler(
		ctx,
		record1.WithID(req.ID, false),
		record1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRecord",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRecordResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteRecord(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteRecord",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRecordResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteRecordResponse{}, nil
}
