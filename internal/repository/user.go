package repository

import (
	"database/sql"
	"errors"
	"plan/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(name, email string) (*model.User, error) {
	query := `
		INSERT INTO users(name, email)
		VALUES($1, $2)
		RETURNING id, name, email, created_at
    `
	user := &model.User{}

	row := r.db.QueryRow(query, name, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}

func (r *UserRepository) GetByID(id int) (*model.User, error) {
	query := `
		SELECT * FROM users WHERE id = $1
    `
	user := &model.User{}

	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}

func (r *UserRepository) Delete(id int) error {
	query := `
		DELETE FROM users WHERE id = $1
    `

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Not found user")
	}
	return nil
}
