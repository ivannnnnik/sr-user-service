package handler

import (
	"context"

	userv1 "github.com/ivannnnnik/sr-proto/gen/go/user/v1"
	"github.com/ivannnnnik/sr-user-service/internal/service"
)

type UserHandler struct{
	userv1.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler{
	return &UserHandler{
		service: svc,
	}
}

func (h *UserHandler) Register(ctx context.Context, req *userv1.RegisterRequest) (*userv1.RegisterResponse, error) {
	user, err := h.service.Register(ctx, req.Email, req.Username, req.Password)
	if err != nil{
		return nil, err
	}

	userConv := UserToProto(user)


	return &userv1.RegisterResponse{
		User: userConv,
	}, nil

}

// func (h *UserHandler) GetProfile(ctx context.Context, req *userv1.GetProfileRequest) (*userv1.GetProfileResponse, error) {
//     // TODO: вызвать h.service.GetProfile
//     // TODO: сконвертировать model.User → userv1.User
//     // TODO: вернуть &userv1.GetProfileResponse{User: ...}
// }
