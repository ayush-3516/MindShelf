package repository

import (
	"context"
	"github.com/ayush-3516/mindshelf/backend/internal/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
	DeleteUser(ctx context.Context, id pgtype.UUID) error
}

type userRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) UserRepository {
	return &userRepository{q: q}
}

func (r *userRepository) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.q.CreateUser(ctx, arg)
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return r.q.GetUserByEmail(ctx, email)
}

func (r *userRepository) GetUserByID(ctx context.Context, id pgtype.UUID) (db.User, error) {
	return r.q.GetUserByID(ctx, id)
}

func (r *userRepository) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return r.q.UpdateUser(ctx, arg)
}

func (r *userRepository) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	return r.q.DeleteUser(ctx, id)
}

