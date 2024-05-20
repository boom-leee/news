package rest

import (
	"encoding/json"
	"net/http"
	inport "news-api/application/port/in"
)

type UserHandlers struct {
	getAllUserUseCase inport.GetAllUsersUseCase
}

func NewUserHandlers(getAllUserUseCase inport.GetAllUsersUseCase) *UserHandlers {
	return &UserHandlers{getAllUserUseCase: getAllUserUseCase}
}

func (u *UserHandlers) GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	usersList, err := u.getAllUserUseCase.Get()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 500,
			Message:    "Unknown err",
		})

	}
	json.NewEncoder(response).Encode(APIResponse[[]*inport.User]{
		StatusCode: 200,
		Message:    "Ok",
		Data:       usersList,
	})
}
