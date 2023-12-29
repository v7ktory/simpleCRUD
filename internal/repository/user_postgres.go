package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/v7ktory/simpleCRUD/internal/model"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) Create(ctx context.Context, user *model.User) (int, error) {
	var id int
	query := "INSERT INTO users (username, email, password, phone) VALUES ($1, $2, $3, $4) RETURNING id"

	row := ur.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.Phone)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UserRepo) GetByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE id = $1"
	err := ur.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepo) Update(ctx context.Context, id int, input *model.UserUpdateInput) error {
	// Проверка существования пользователя
	exists, err := ur.UserExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user does not exist")
	}

	// Запрос на обновление
	query := "UPDATE users SET username = $1, phone = $2 WHERE id = $3"
	result, err := ur.db.ExecContext(ctx, query, input.Username, input.Phone, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	return nil
}

func (ur *UserRepo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := ur.db.ExecContext(ctx, query, id)
	return err
}

func (ur *UserRepo) UserExists(ctx context.Context, userID int) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)"
	var exists bool
	err := ur.db.GetContext(ctx, &exists, query, userID)
	return exists, err
}
