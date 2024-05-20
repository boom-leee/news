package outAdapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	db2 "news-api/postgres/db"
)

type CategoryAdapter struct {
	pool *pgxpool.Pool
}

func NewCategoryAdapter(pool *pgxpool.Pool) *CategoryAdapter {
	return &CategoryAdapter{pool: pool}
}

func (u *CategoryAdapter) GetAll() ([]db2.Category, error) {
	query := db2.New(u.pool)

	return query.GetAllCategories(context.Background())

}
