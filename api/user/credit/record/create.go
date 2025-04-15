package record

import (
	"context"

	record1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/user/credit/record"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/credit/record"
)

func (s *Server) CreateRecord(ctx context.Context, in *npool.CreateRecordRequest) (*npool.CreateRecordResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateRecord",
			"In", in,
		)
		return &npool.CreateRecordResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := record1.NewHandler(
		ctx,
		record1.WithEntID(req.EntID, false),
		record1.WithAppID(req.AppID, true),
		record1.WithUserID(req.UserID, true),
		record1.WithOperationType(req.OperationType, true),
		record1.WithCreditsChange(req.CreditsChange, true),
		record1.WithExtra(req.Extra, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRecord",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRecordResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateRecord(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateRecord",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRecordResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateRecordResponse{}, nil
}
