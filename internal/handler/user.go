package handler

import (
	"context"

	userv1 "github.com/ivannnnnik/sr-proto/gen/go/user/v1"
	"github.com/ivannnnnik/sr-user-service/internal/model"
)

type userService interface {
    Register(ctx context.Context, email, username, password string) (*model.User, error)
    GetProfile(ctx context.Context, id string) (*model.User, error)
}

type UserHandler struct{
	userv1.UnimplementedUserServiceServer
	service userService
}

func NewUserHandler(svc userService) *UserHandler{
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

func (h *UserHandler) GetProfile(ctx context.Context, req *userv1.GetProfileRequest) (*userv1.GetProfileResponse, error) {
	user, err := h.service.GetProfile(ctx, req.UserId)
	if err != nil{
		return nil, err
	}

	userConv := UserToProto(user)

	return &userv1.GetProfileResponse{
		User: userConv,
	}, nil

}
