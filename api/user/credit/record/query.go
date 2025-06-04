package record

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	record1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/user/credit/record"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/credit/record"
)

func (s *Server) GetRecord(ctx context.Context, in *npool.GetRecordRequest) (*npool.GetRecordResponse, error) {
	handler, err := record1.NewHandler(
		ctx,
		record1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecord",
			"In", in,
			"Error", err,
		)
		return &npool.GetRecordResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetRecord(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecord",
			"In", in,
			"Error", err,
		)
		return &npool.GetRecordResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRecordResponse{
		Info: info,
	}, nil
}

func (s *Server) GetRecords(ctx context.Context, in *npool.GetRecordsRequest) (*npool.GetRecordsResponse, error) {
	handler, err := record1.NewHandler(
		ctx,
		record1.WithConds(in.GetConds()),
		record1.WithOffset(in.GetOffset()),
		record1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecords",
			"In", in,
			"Error", err,
		)
		return &npool.GetRecordsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.GetRecords(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecords",
			"In", in,
			"Error", err,
		)
		return &npool.GetRecordsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRecordsResponse{
		Infos: infos,
	}, nil
}
