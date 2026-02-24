package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"plan/old/internal/model"
	"strings"
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

func (r *UserRepository) Update(id int, user *model.UpdateUser) (*model.User, error) {
	paramsForQuery := []string{}
	valuesForQuery := []any{}
	index := 1

	if user.Name != nil {
		paramsForQuery = append(paramsForQuery, fmt.Sprintf("name = $%d", index))
		valuesForQuery = append(valuesForQuery, *user.Name)
		index++
	}

	if user.Email != nil {
		paramsForQuery = append(paramsForQuery, fmt.Sprintf("email = $%d", index))
		valuesForQuery = append(valuesForQuery, *user.Email)
		index++
	}

	args := append(valuesForQuery, id)
	query := fmt.Sprintf(`UPDATE users SET %s WHERE id = $%d RETURNING id, name, email, created_at`,
		strings.Join(paramsForQuery, ", "),
		index,
	)

	userM := &model.User{}

	err := r.db.QueryRow(query, args...).Scan(&userM.ID, &userM.Name, &userM.Email, &userM.CreatedAt)
	return userM, err
}
