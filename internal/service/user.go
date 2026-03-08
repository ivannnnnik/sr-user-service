package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/ivannnnnik/sr-user-service/internal/model"
)


type userRepo interface {
	Create(ctx context.Context, user *model.User) error
    GetByID(ctx context.Context, id string) (*model.User, error)
}

type UserService struct{
	repo userRepo
}

func NewUserService(repo userRepo) *UserService{
	return &UserService{
		repo: repo,
	}
}


func (svc *UserService) Register(ctx context.Context, email, username, password string) (*model.User, error){
	
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return nil, err
	}
	
	userModel := model.User{
		Email: email,
		Username: username,
		PasswordHash: string(hash),
	}
	
	err = svc.repo.Create(ctx, &userModel)

	if err != nil{
		return nil, err
	}

	return &userModel, nil

}


func (svc *UserService) GetProfile(ctx context.Context, id string) (*model.User, error){
	user, err := svc.repo.GetByID(ctx, id)
	if err != nil{
		return nil, err
	}

	return user, nil

}