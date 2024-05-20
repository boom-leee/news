package service

import (
	inport "news-api/application/port/in"
	outport "news-api/application/port/out"
)

type CategoriesService struct {
	categoriesPort outport.Categories
}

func NewCategoriesService(categoriesPort outport.Categories) *CategoriesService {
	return &CategoriesService{categoriesPort: categoriesPort}
}

func (g *CategoriesService) GetAll() ([]*inport.Category, error) {
	categoriesList, err := g.categoriesPort.GetAll()
	if err != nil {
		return nil, err
	}
	return func() []*inport.Category {
		result := make([]*inport.Category, len(categoriesList))
		for i, v := range categoriesList {
			result[i] = MapCategory(v)
		}
		return result
	}(), nil
}
