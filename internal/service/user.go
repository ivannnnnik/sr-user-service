package service

import (
	"context"
	"os/user"

	"golang.org/x/crypto/bcrypt"

	"github.com/ivannnnnik/sr-user-service/internal/model"
	"github.com/ivannnnnik/sr-user-service/internal/repository"
)

type UserService struct{
	repo *repository.UserRepository
}

func NewUserRepository(repo *repository.UserRepository) *UserService{
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