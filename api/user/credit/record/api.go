package record

import (
	"github.com/NpoolPlatform/message/npool/billing/mw/v1/user/credit/record"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	record.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	record.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
