package usecase

import (
	"plan/old/internal/model"
	"plan/old/internal/repository"
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

func (u *User) Update(id int, user *model.UpdateUser) (*model.User, error) {
	_, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	userUp, err := u.userRepo.Update(id, user)
	if err != nil {
		return nil, err
	}

	return userUp, nil
}
