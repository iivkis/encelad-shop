package httphandler

import "github.com/go-chi/chi/v5"

type HTTPHandler struct {
	Router          chi.Router
	UserHTTPHandler *UserHTTPHandler
}

func NewHTTPHandler(
	UserHTTPHandler *UserHTTPHandler,
) *HTTPHandler {
	handler := &HTTPHandler{
		Router:          chi.NewRouter(),
		UserHTTPHandler: UserHTTPHandler,
	}
	handler.setupAPI()
	return handler
}

func (h *HTTPHandler) setupAPI() {
	api := chi.NewRouter()

	api.Route("/users", func(r chi.Router) {
		r.Get("/{id}", h.UserHTTPHandler.GetByID)
		r.Post("/", h.UserHTTPHandler.Create)
		r.Put("/{id}", h.UserHTTPHandler.Update)
	})

	h.Router.Mount("/api/v1", api)
}
