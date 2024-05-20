package outAdapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	db2 "news-api/postgres/db"
)

type UserAdapter struct {
	pool *pgxpool.Pool
}

func NewUserAdapter(pool *pgxpool.Pool) *UserAdapter {
	return &UserAdapter{pool: pool}
}

func (u *UserAdapter) Get() ([]db2.User, error) {
	query := db2.New(u.pool)

	return query.GetAllUsers(context.Background())

}
