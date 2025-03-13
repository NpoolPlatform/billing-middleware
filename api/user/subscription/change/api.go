package change

import (
	"github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription/change"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	change.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	change.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
