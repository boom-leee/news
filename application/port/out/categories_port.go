package outport

import "news-api/internal/db"

type Categories interface {
	GetAll() ([]db.Category, error)
}
