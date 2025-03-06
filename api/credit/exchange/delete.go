package exchange

import (
	"context"

	exchange1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/credit/exchange"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
)

func (s *Server) DeleteExchange(ctx context.Context, in *npool.DeleteExchangeRequest) (*npool.DeleteExchangeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteExchange",
			"In", in,
		)
		return &npool.DeleteExchangeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithID(req.ID, false),
		exchange1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteExchange",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteExchange(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteExchange",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteExchangeResponse{}, nil
}
