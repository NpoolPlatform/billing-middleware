package addon

import (
	"github.com/NpoolPlatform/message/npool/billing/mw/v1/addon"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	addon.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	addon.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
