package exchange

import (
	"context"

	exchange1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/credit/exchange"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
)

func (s *Server) UpdateExchange(ctx context.Context, in *npool.UpdateExchangeRequest) (*npool.UpdateExchangeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateExchange",
			"In", in,
		)
		return &npool.UpdateExchangeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithID(req.ID, false),
		exchange1.WithEntID(req.EntID, false),
		exchange1.WithCredit(req.Credit, false),
		exchange1.WithExchangeThreshold(req.ExchangeThreshold, false),
		exchange1.WithPath(req.Path, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateExchange",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateExchange(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateExchange",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateExchangeResponse{}, nil
}
