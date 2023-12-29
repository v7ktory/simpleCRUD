package service

import (
	"context"

	"github.com/v7ktory/simpleCRUD/internal/model"
	"github.com/v7ktory/simpleCRUD/internal/repository"
)

type UserService struct {
	repos repository.User
}

func NewUserService(repos repository.User) *UserService {
	return &UserService{
		repos: repos,
	}
}

func (us *UserService) Create(ctx context.Context, user *model.User) (int, error) {
	return us.repos.Create(ctx, user)
}

func (us *UserService) GetByID(ctx context.Context, id int) (*model.User, error) {
	return us.repos.GetByID(ctx, id)
}

func (us *UserService) Update(ctx context.Context, id int, input *model.UserUpdateInput) error {
	return us.repos.Update(ctx, id, input)
}

func (us *UserService) Delete(ctx context.Context, id int) error {
	return us.repos.Delete(ctx, id)
}
