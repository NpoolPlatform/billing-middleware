package exchange

import (
	"github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	exchange.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	exchange.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
