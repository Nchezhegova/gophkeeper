package grpc

import (
	"github.com/Nchezhegova/gophkeeper/internal/interfaces/grpc/proto"
	"github.com/Nchezhegova/gophkeeper/internal/usecases"
)

type GophKeeperServer struct {
	proto.UnimplementedGophKeeperServer
	UserUseCase usecases.UserUseCase
	DataUseCase usecases.DataUseCase
}
