package service

import (
	"context"

	"github.com/v7ktory/simpleCRUD/internal/model"
	"github.com/v7ktory/simpleCRUD/internal/repository"
)

type User interface {
	Create(ctx context.Context, user *model.User) (int, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context, id int, input *model.UserUpdateInput) error
	Delete(ctx context.Context, id int) error
}
type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
