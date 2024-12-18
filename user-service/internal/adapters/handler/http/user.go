package httphandler

import (
	"enceland_user-service/internal/core/ports"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHTTPHandler struct {
	UserSvc ports.UserService
}

func NewUserHTTPHandler(UserService ports.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		UserSvc: UserService,
	}
}

type UserGetByIDResponse struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
}

func (h *UserHTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	user, err := h.UserSvc.GetByID(r.Context(), id)
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(UserGetByIDResponse{
		ID:        user.GetID(),
		Firstname: user.GetFirstname(),
		Lastname:  user.GetLastname(),
		Role:      user.GetRole().String(),
	}); err != nil {
		panic(err)
	}
}
