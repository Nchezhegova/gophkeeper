package grpc

import (
	"context"
	"github.com/Nchezhegova/gophkeeper/internal/entities"
	"github.com/Nchezhegova/gophkeeper/internal/interfaces/grpc/proto"
	"github.com/Nchezhegova/gophkeeper/internal/usecases"
)

type Server struct {
	proto.UnimplementedGophKeeperServer
	UserUseCase usecases.UserUseCase
	DataUseCase usecases.DataUseCase
}

func (s *Server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	err := s.UserUseCase.Register(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.RegisterResponse{}, nil
}

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	token, err := s.UserUseCase.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.LoginResponse{Token: token}, nil
}

func (s *Server) StoreData(ctx context.Context, req *proto.StoreDataRequest) (*proto.StoreDataResponse, error) {
	userID := uint32(ctx.Value("user_id").(int))
	err := s.DataUseCase.StoreData(userID, req.Key, req.Type, req.Data)
	if err != nil {
		return nil, err
	}
	return &proto.StoreDataResponse{}, nil
}

func (s *Server) GetData(ctx context.Context, req *proto.GetDataRequest) (*proto.GetDataResponse, error) {
	userID := uint32(ctx.Value("user_id").(int))
	var data []entities.Data
	var err error

	data, err = s.DataUseCase.GetData(userID, req.Key, req.Type, req.Identifier)

	if err != nil {
		return nil, err
	}

	var protoData []*proto.Data
	for _, d := range data {
		protoData = append(protoData, &proto.Data{
			Id:   uint32(d.ID),
			Type: d.Type,
			Data: d.Data,
			Key:  d.Key,
		})
	}
	return &proto.GetDataResponse{Data: protoData}, nil
}

func (s *Server) UpdateData(ctx context.Context, req *proto.UpdateDataRequest) (*proto.UpdateDataResponse, error) {
	userID := uint32(ctx.Value("user_id").(int))
	err := s.DataUseCase.UpdateData(userID, req.Key, req.Type, req.Identifier, req.NewData)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateDataResponse{}, nil
}

func (s *Server) DeleteData(ctx context.Context, req *proto.DeleteDataRequest) (*proto.DeleteDataResponse, error) {
	userID := uint32(ctx.Value("user_id").(int))
	err := s.DataUseCase.DeleteData(userID, req.Key, req.Type, req.Identifier)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteDataResponse{}, nil
}
