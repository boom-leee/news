package outport

import (
	"news-api/postgres/db"
)

type GetAllUsers interface {
	Get() ([]db.User, error)
}
