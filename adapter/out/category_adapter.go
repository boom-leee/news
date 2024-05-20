package outAdapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"news-api/internal/db"
)

type CategoryAdapter struct {
	pool *pgxpool.Pool
}

func NewCategoryAdapter(pool *pgxpool.Pool) *CategoryAdapter {
	return &CategoryAdapter{pool: pool}
}

func (u *CategoryAdapter) GetAll() ([]db.Category, error) {
	query := db.New(u.pool)

	return query.GetAllCategories(context.Background())

}
