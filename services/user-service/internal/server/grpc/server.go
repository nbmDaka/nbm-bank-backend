package grpc

import (
	"log"
	"net"

	pb "github.com/nbmDaka/nbm-bank-backend/services/user-service/proto/user"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/auth"

	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server

	port string
}

func NewServer(
	port string,
	userHandler pb.UserServiceServer,
) *Server {

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			auth.UnaryAuthInterceptor,
		),
	)

	pb.RegisterUserServiceServer(
		grpcServer,
		userHandler,
	)

	return &Server{
		server: grpcServer,
		port:   port,
	}
}

func (s *Server) Start() error {

	log.Println(
		"gRPC server starting on",
		s.port,
	)

	listener, err := net.Listen(
		"tcp",
		":"+s.port,
	)

	if err != nil {
		return err
	}

	return s.server.Serve(listener)
}

func (s *Server) Shutdown() {

	s.server.GracefulStop()

}
