package v1

import (
	"github.com/go-chi/chi/v5"
)

func Router(h *ApiHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/products", h.GetProducts)

	return r
}
