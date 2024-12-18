package http

import (
	v1 "github.com/Kroning/mytheresa/internal/transport/http/v1"
	"github.com/go-chi/chi/v5"
)

func NewRouter(
	httpApiHandlerV1 *v1.ApiHandler,
) chi.Router {
	r := chi.NewRouter()

	r.Get("/status", routerCheck)

	r.Route("/api/v1", func(r chi.Router) {
		// TODO: metrics middleware
		// TODO: auth middleware
		r.Mount("/", v1.Router(httpApiHandlerV1))
	})

	return r
}
