package rest

import (
	"encoding/json"
	"net/http"
	inport "news-api/application/port/in"
)

type DummyHandler struct {
	dummyUseCase inport.Dummy
}

func NewDummyHandler(dummyUseCase inport.Dummy) *DummyHandler {
	return &DummyHandler{dummyUseCase: dummyUseCase}
}

func (d *DummyHandler) Dummy(response http.ResponseWriter, request *http.Request) {
	dummyOutput := d.dummyUseCase.Get()
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(dummyOutput)
}
