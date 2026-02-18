package usecase

import (
	"plan/internal/model"
	"plan/internal/repository"
)

type User struct {
	userRepo *repository.UserRepository
}

func NewUser(repo *repository.UserRepository) *User {
	return &User{
		userRepo: repo,
	}
}

func (u *User) Create(name, email string) (*model.User, error) {
	user, err := u.userRepo.Create(name, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) GetByID(id int) (*model.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Delete(id int) error {
	_, err := u.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	err = u.userRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
