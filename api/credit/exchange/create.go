package exchange

import (
	"context"

	exchange1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/credit/exchange"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
)

func (s *Server) CreateExchange(ctx context.Context, in *npool.CreateExchangeRequest) (*npool.CreateExchangeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateExchange",
			"In", in,
		)
		return &npool.CreateExchangeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithEntID(req.EntID, false),
		exchange1.WithAppID(req.AppID, true),
		exchange1.WithUsageType(req.UsageType, true),
		exchange1.WithCredit(req.Credit, true),
		exchange1.WithExchangeThreshold(req.ExchangeThreshold, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateExchange",
			"In", in,
			"Error", err,
		)
		return &npool.CreateExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateExchange(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateExchange",
			"In", in,
			"Error", err,
		)
		return &npool.CreateExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateExchangeResponse{}, nil
}
