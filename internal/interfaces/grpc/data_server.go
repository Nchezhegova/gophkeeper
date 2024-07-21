package grpc

import (
	"context"
	"github.com/Nchezhegova/gophkeeper/internal/entities"
	"github.com/Nchezhegova/gophkeeper/internal/interfaces/grpc/proto"
	"github.com/Nchezhegova/gophkeeper/internal/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GophKeeperServer) StoreData(ctx context.Context, req *proto.StoreDataRequest) (*proto.StoreDataResponse, error) {
	userID, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "user ID not found in context")
	}
	err := s.DataUseCase.StoreData(userID, req.Key, req.Type, req.Data)
	if err != nil {
		return nil, err
	}
	return &proto.StoreDataResponse{}, nil
}

func (s *GophKeeperServer) GetData(ctx context.Context, req *proto.GetDataRequest) (*proto.GetDataResponse, error) {
	userID, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "user ID not found in context")
	}

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

func (s *GophKeeperServer) UpdateData(ctx context.Context, req *proto.UpdateDataRequest) (*proto.UpdateDataResponse, error) {
	userID, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "user ID not found in context")
	}
	err := s.DataUseCase.UpdateData(userID, req.Key, req.Type, req.Identifier, req.NewData)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateDataResponse{}, nil
}

func (s *GophKeeperServer) DeleteData(ctx context.Context, req *proto.DeleteDataRequest) (*proto.DeleteDataResponse, error) {
	userID, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "user ID not found in context")
	}
	err := s.DataUseCase.DeleteData(userID, req.Key, req.Type, req.Identifier)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteDataResponse{}, nil
}
