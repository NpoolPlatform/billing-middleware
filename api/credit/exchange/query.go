package exchange

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	exchange1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/credit/exchange"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
)

func (s *Server) GetExchange(ctx context.Context, in *npool.GetExchangeRequest) (*npool.GetExchangeResponse, error) {
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetExchange",
			"In", in,
			"Error", err,
		)
		return &npool.GetExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetExchange(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetExchange",
			"In", in,
			"Error", err,
		)
		return &npool.GetExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetExchangeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetExchanges(ctx context.Context, in *npool.GetExchangesRequest) (*npool.GetExchangesResponse, error) {
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithConds(in.GetConds()),
		exchange1.WithOffset(in.GetOffset()),
		exchange1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetExchanges",
			"In", in,
			"Error", err,
		)
		return &npool.GetExchangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.GetExchanges(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetExchanges",
			"In", in,
			"Error", err,
		)
		return &npool.GetExchangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetExchangesResponse{
		Infos: infos,
	}, nil
}
