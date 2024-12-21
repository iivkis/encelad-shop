package httphandler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HttpHandler struct {
	router          chi.Router
	userHttpHandler *UserHttpHandler
}

func NewHttpHandler(
	userHttpHandler *UserHttpHandler,
) http.Handler {
	handler := &HttpHandler{
		router:          chi.NewRouter(),
		userHttpHandler: userHttpHandler,
	}
	handler.setup()
	return handler.router
}

func (h *HttpHandler) setup() {
	api := chi.NewRouter()

	api.Route("/users", func(r chi.Router) {
		r.Get("/{id}", h.userHttpHandler.GetByID)
		r.Post("/", h.userHttpHandler.Create)
		r.Put("/{id}", h.userHttpHandler.Update)
	})

	h.router.Mount("/api/v1", api)
}
