package routes

import (
	"github.com/ramadhanalfarisi/go-codebase/services/user/grpc"
	"github.com/ramadhanalfarisi/go-codebase/services/user/usecase"
)

func UserGrpcRoute() grpc.UserControllerServer {
	userUsecase := usecase.NewUsecaseGrpc()
	return userUsecase
}
