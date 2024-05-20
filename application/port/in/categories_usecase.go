package inport

import "time"

type Category struct {
	ID        string
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type CategoriesUseCase interface {
	GetAll() ([]*Category, error)
}
