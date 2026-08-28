package grpc

import (
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	port string
}

func NewServer(port string) *Server {

	return &Server{
		server: grpc.NewServer(),
		port: port,
	}

}

func (s *Server) Start() error {

	println("GRPC server starting on " + s.port)

	listener, err := net.Listen(
		"tcp",
		":" + s.port,
	)

	if err != nil {
		return err
	}

	return s.server.Serve(listener)

}

func (s *Server) Shutdown() {

	s.server.GracefulStop()

}
