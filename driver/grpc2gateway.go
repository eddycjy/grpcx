package driver

import (
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

type GRPC2GatewayIface interface {
	DriverIface
	GetEngine() *runtime.ServeMux
	GetHTTPMux() *http.ServeMux
}

type gRPC2GatewayServer struct {
	grpcS        *grpc.Server
	grpcGatewayS *runtime.ServeMux
	httpS        *http.ServeMux
	lis          net.Listener
}

func NewGRPC2GatewayServer(gRPCS *grpc.Server, httpS *http.ServeMux, grpcGatewayS *runtime.ServeMux) GRPC2GatewayIface {
	return &gRPC2GatewayServer{grpcS: gRPCS, httpS: httpS, grpcGatewayS: grpcGatewayS}
}

func (s *gRPC2GatewayServer) GetEngine() *runtime.ServeMux {
	return s.grpcGatewayS
}

func (s *gRPC2GatewayServer) GetHTTPMux() *http.ServeMux {
	return s.httpS
}

func (s *gRPC2GatewayServer) SetListener(lis net.Listener) {
	s.lis = lis
}

func (s *gRPC2GatewayServer) Serve() error {
	serveMux := s.GetHTTPMux()
	serveMux.Handle("/", s.GetEngine())
	return http.Serve(s.lis, s.grpcHandlerFunc(s.grpcS, serveMux))
}

func (s *gRPC2GatewayServer) grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
