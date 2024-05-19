package rest

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func AppRouter(dummyHandler *DummyHandler) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	router.Get("/dummy", dummyHandler.Dummy)
	return router
}
