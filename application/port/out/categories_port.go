package outport

import (
	"news-api/postgres/db"
)

type Categories interface {
	GetAll() ([]db.Category, error)
}
