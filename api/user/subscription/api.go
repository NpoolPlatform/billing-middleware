package subscription

import (
	"github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	subscription.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	subscription.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
