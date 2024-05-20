package rest

import (
	"encoding/json"
	"net/http"
	inport "news-api/application/port/in"
)

type CategoryHandlers struct {
	categoryUseCase inport.CategoriesUseCase
}

func NewCategoryHandlers(categoryUseCase inport.CategoriesUseCase) *CategoryHandlers {
	return &CategoryHandlers{categoryUseCase: categoryUseCase}
}

func (u *CategoryHandlers) GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	categoriesList, err := u.categoryUseCase.GetAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 500,
			Message:    "Unknown err",
		})

	}
	json.NewEncoder(response).Encode(APIResponse[[]*inport.Category]{
		StatusCode: 200,
		Message:    "Ok",
		Data:       categoriesList,
	})
}
