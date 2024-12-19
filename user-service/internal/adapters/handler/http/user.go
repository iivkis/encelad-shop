package httphandler

import (
	"encelad-shared/core/ports"
	"encelad-shared/pkg/responder"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHTTPHandler struct {
	UserService ports.UserService
	responder   *responder.Responder
}

func NewUserHTTPHandler(UserService ports.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		UserService: UserService,
		responder:   responder.NewResponder(),
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
		h.responder.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.UserService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ports.ErrNotFound) {
			h.responder.JsonErr(w, http.StatusNotFound, fmt.Errorf("user not found with id: %d", id))
		} else {
			h.responder.JsonErr(w, http.StatusInternalServerError, nil)
		}
		return
	}

	h.responder.JsonOk(
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

type UserCreateRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type UserCreateResponse struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
}

func (h *UserHTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		h.responder.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	userIn := ports.CreateUserIn(body)

	user, err := h.UserService.Create(r.Context(), &userIn)
	if err != nil {
		h.responder.JsonErr(w, http.StatusInternalServerError, nil)
		return
	}

	h.responder.JsonOk(
		w,
		http.StatusCreated,
		UserCreateResponse{
			ID:        user.GetID(),
			Firstname: user.GetFirstname(),
			Lastname:  user.GetLastname(),
			Role:      user.GetRole().String(),
		},
	)
}

type UserUpdateRequest struct {
	UserCreateRequest
}

type UserUpdateResponse struct {
	UserCreateResponse
}

func (h *UserHTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.responder.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	var body UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		h.responder.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	userIn := ports.UpdateUserIn{
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
	}

	user, err := h.UserService.Update(r.Context(), id, &userIn)
	if err != nil {
		if errors.Is(err, ports.ErrNotFound) {
			h.responder.JsonErr(w, http.StatusNotFound, fmt.Errorf("user not found with id: %d", id))
		} else {
			h.responder.JsonErr(w, http.StatusInternalServerError, nil)
		}
		return
	}

	h.responder.JsonOk(
		w,
		http.StatusOK,
		UserUpdateResponse{
			UserCreateResponse{
				ID:        user.GetID(),
				Firstname: user.GetFirstname(),
				Lastname:  user.GetLastname(),
				Role:      user.GetRole().String(),
			},
		},
	)
}
