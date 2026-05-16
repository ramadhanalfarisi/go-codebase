package usecase

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/ramadhanalfarisi/go-codebase/constants"
	"github.com/ramadhanalfarisi/go-codebase/helpers"
	"github.com/ramadhanalfarisi/go-codebase/services/user/grpc"
)

type UsecaseGrpc struct {
	grpc.UnimplementedUserControllerServer

	mu sync.RWMutex
}

func NewUsecaseGrpc() grpc.UserControllerServer {
	return &UsecaseGrpc{}
}

func (u *UsecaseGrpc) Middleware(ctx context.Context, input *grpc.MiddlewareInput) (*grpc.UserDetail, error) {
	if input == nil {
		err := fmt.Errorf("input have to filled")
		helpers.Error(err)
		return nil, err
	}
	userDetail, err := u.tokenVerification(input.Token)
	if err != nil {
		return nil, err
	}
	if userDetail.Id <= 0 {
		return nil, fmt.Errorf("user not found")
	}
	return &grpc.UserDetail{
		Id:    float32(userDetail.Id),
		Email: userDetail.Email,
		Roles: userDetail.Roles,
	}, nil
}

func (u *UsecaseGrpc) tokenVerification(token string) (*helpers.UserDetail, error) {
	if !strings.Contains(token, "Bearer") {
		helpers.Error(fmt.Errorf("Have to a Bearer token"))
		return nil, fmt.Errorf(constants.InvalidToken)
	} else {
		tokenString := strings.Replace(token, "Bearer ", "", -1)
		claims, err := helpers.ParseUserJWT(tokenString)
		if err != nil {
			helpers.Error(err)
			return nil, fmt.Errorf(constants.InvalidToken)
		}
		userDetail := helpers.GetUserDetail(claims)
		return &userDetail, nil
	}
}
