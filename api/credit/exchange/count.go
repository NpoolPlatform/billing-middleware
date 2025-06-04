//nolint:dupl
package exchange

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	exchange1 "github.com/NpoolPlatform/billing-middleware/pkg/mw/credit/exchange"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
)

func (s *Server) CountExchanges(ctx context.Context, in *npool.CountExchangesRequest) (*npool.CountExchangesResponse, error) {
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountExchanges",
			"In", in,
			"Error", err,
		)
		return &npool.CountExchangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	total, err := handler.CountExchanges(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountExchanges",
			"In", in,
			"Error", err,
		)
		return &npool.CountExchangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CountExchangesResponse{
		Total: total,
	}, nil
}
