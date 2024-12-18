package httphandler

import (
	"enceland_user-service/internal/core/ports"
	"errors"
	"fmt"
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
		HTTPJsonErr(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.UserSvc.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ports.ErrNotFound) {
			HTTPJsonErr(w, http.StatusNotFound, fmt.Errorf("user not found with id: %d", id))
		} else {
			HTTPJsonErr(w, http.StatusInternalServerError, nil)
		}
		return
	}

	HTTPJsonOk(
		w,
		http.StatusOK,
		UserGetByIDResponse{
			ID:        user.GetID(),
			Firstname: user.GetFirstname(),
			Lastname:  user.GetLastname(),
			Role:      user.GetRole().String(),
		},
	)
}
