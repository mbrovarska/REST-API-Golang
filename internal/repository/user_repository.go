package repository

import (
	"context"
	"fmt"
	"time"

	models "example.com/rest-api-notes/internal/domain"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
	qb sq.StatementBuilderType
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *UserRepository) SignUp(ctx context.Context, email, hashedPassword string) (*models.User, error) {
	created_at := time.Now()
	updated_at := time.Now()
	query, args, err := r.qb.Insert("users").Columns("email", "password", "created_at", "updated_at").
		Values(email, hashedPassword, created_at, updated_at).
		Suffix("RETURNING id, email, created_at, updated_at").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var user models.User
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil

}

func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	query, args, err := r.qb.Select("COUNT(*)").From("users").Where(sq.Eq{"email": email}).
		ToSql()

	if err != nil {
		return false, fmt.Errorf("failed to build query: %w", err)
	}

	var count int
	err = r.db.QueryRow(ctx, query, args...).Scan(&count)

	if err != nil {
		return false, fmt.Errorf("failed to check email: %w", err)
	}

	return count > 0, nil 
}