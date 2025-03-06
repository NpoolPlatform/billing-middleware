package exchange

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	exchange1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/credit/exchange"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
)

func (s *Server) ExistExchangeConds(ctx context.Context, in *npool.ExistExchangeCondsRequest) (*npool.ExistExchangeCondsResponse, error) {
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistExchangeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistExchangeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistExchangeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistExchangeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistExchangeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistExchangeCondsResponse{
		Info: exist,
	}, nil
}
