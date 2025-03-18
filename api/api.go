package api

import (
	"context"

	addon "github.com/NpoolPlatform/billing-middleware/api/addon"
	exchange "github.com/NpoolPlatform/billing-middleware/api/credit/exchange"
	subscription "github.com/NpoolPlatform/billing-middleware/api/subscription"
	record "github.com/NpoolPlatform/billing-middleware/api/user/credit/record"
	usersubscription "github.com/NpoolPlatform/billing-middleware/api/user/subscription"
	usersubscriptionchange "github.com/NpoolPlatform/billing-middleware/api/user/subscription/change"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterMiddlewareServer(server, &Server{})
	addon.Register(server)
	subscription.Register(server)
	exchange.Register(server)
	usersubscription.Register(server)
	record.Register(server)
	usersubscriptionchange.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := npool.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := addon.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := subscription.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := exchange.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := usersubscription.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := record.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := usersubscriptionchange.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
