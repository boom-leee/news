package outAdapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"news-api/internal/db"
)

type UserAdapter struct {
	pool *pgxpool.Pool
}

func NewUserAdapter(pool *pgxpool.Pool) *UserAdapter {
	return &UserAdapter{pool: pool}
}

func (u *UserAdapter) Get() ([]db.User, error) {
	query := db.New(u.pool)

	return query.GetAllUsers(context.Background())

}
