package repository

import (
	"context"

	"github.com/ivannnnnik/sr-user-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct{
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository{
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error{
		
	query := ` INSERT INTO users(email, password_hash)
			    VALUES ($1, $2)
				RETURNING id, email, created_at`
	
	err := r.db.QueryRowContext(ctx, query, user.Email, user.PasswordHash).
	Scan(&user.ID, &user.Email,&user.CreatedAt)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*model.User, error){
	query := `SELECT id, email, created_at FROM users WHERE id = $1`

	var user model.User
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil{
		return nil, err
	}

	return &user, nil
}