package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/v7ktory/simpleCRUD/internal/model"
)

type User interface {
	Create(ctx context.Context, user *model.User) (int, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context, id int, input *model.UserUpdateInput) error
	Delete(ctx context.Context, id int) error
}
type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepo(db),
	}
}
