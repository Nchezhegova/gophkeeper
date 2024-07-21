package main

import (
	servergrpc "github.com/Nchezhegova/gophkeeper/internal/interfaces/grpc"
	"github.com/Nchezhegova/gophkeeper/internal/interfaces/grpc/proto"
	"github.com/Nchezhegova/gophkeeper/internal/interfaces/repository"
	"github.com/Nchezhegova/gophkeeper/internal/middleware"
	"github.com/Nchezhegova/gophkeeper/internal/tlsconfig"
	"github.com/Nchezhegova/gophkeeper/internal/usecases"
	"github.com/Nchezhegova/gophkeeper/pkg/db"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	database := db.New()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tlsCredentials, err := tlsconfig.LoadServerTLSCredentials()
	if err != nil {
		log.Fatalf("failed to load TLS credentials: %v", err)
	}

	serverOptions := []grpc.ServerOption{
		grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(middleware.UnaryServerInterceptor()),
	}
	s := grpc.NewServer(serverOptions...)

	userRepository := repository.UserRepository{DB: database}
	userUseCase := usecases.UserUseCase{UserRepository: userRepository}

	dataRepository := repository.DataRepository{DB: database}
	dataUseCase := usecases.DataUseCase{DataRepository: dataRepository}

	server := &servergrpc.GophKeeperServer{
		UserUseCase: userUseCase,
		DataUseCase: dataUseCase,
	}

	proto.RegisterGophKeeperServer(s, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
