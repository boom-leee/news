package outport

import "news-api/internal/db"

type GetAllUsers interface {
	Get() ([]db.User, error)
}
