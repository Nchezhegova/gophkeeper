package grpc

import (
	"context"
	"github.com/Nchezhegova/gophkeeper/internal/interfaces/grpc/proto"
)

func (s *GophKeeperServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	err := s.UserUseCase.Register(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.RegisterResponse{}, nil
}

func (s *GophKeeperServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	token, err := s.UserUseCase.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.LoginResponse{Token: token}, nil
}
