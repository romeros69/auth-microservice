package rpc

import (
	"auth-microservice/config"
	"auth-microservice/internal/controller/grpc/gen"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ServerRpc struct {
	l net.Listener

	server *grpc.Server
}

func New(cfg *config.Config) (*ServerRpc, error) {
	listen, err := net.Listen(cfg.GrpcProtocol, cfg.GrpcURL)
	if err != nil {
		return nil, fmt.Errorf("error in init grpc - lisening: %w", err)
	}

	return &ServerRpc{
		l:      listen,
		server: grpc.NewServer(),
	}, nil
}

func (r *ServerRpc) Start(authRegSrv gen.AuthRegServer) error {
	go func() {
		gen.RegisterAuthRegServer(r.server, authRegSrv)
		if err := r.server.Serve(r.l); err != nil {
			log.Fatalf("Cannot start grpc server: %v", err)
		}
	}()

	log.Printf("gRpc successfully started on: %s", r.l.Addr())
	return nil
}
