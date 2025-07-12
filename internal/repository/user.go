package repository

import (
	"context"

	"github.com/ymatsukawa/backbase-go/internal/entity"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type userRepository struct {
	queries *entity.Queries
}

func NewUserRepository(queries *entity.Queries) UserRepository {
	return &userRepository{
		queries: queries,
	}
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return r.queries.GetAllUsers(ctx)
}
