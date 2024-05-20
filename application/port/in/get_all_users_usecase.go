package inport

import "time"

type User struct {
	ID        string
	AuthID    string
	Email     string
	Name      string
	Role      string
	ImageUrl  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type GetAllUsersUseCase interface {
	Get() ([]*User, error)
}
