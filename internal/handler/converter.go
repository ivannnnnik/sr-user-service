package handler

import (
	"time"

	userv1 "github.com/ivannnnnik/sr-proto/gen/go/user/v1"
	"github.com/ivannnnnik/sr-user-service/internal/model"
)

func UserToProto(user *model.User) *userv1.User{
	return &userv1.User{
		Id: user.ID,
		Email: user.Email,
		Username: user.Username,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}