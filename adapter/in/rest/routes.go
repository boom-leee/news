package rest

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func AppRouter(dummyHandler *DummyHandler, userHandlers *UserHandlers, categoryHandlers *CategoryHandlers) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	router.Get("/dummy", dummyHandler.Dummy)
	router.Get("/users", userHandlers.GetAll)
	router.Get("/categories", categoryHandlers.GetAll)
	return router
}
