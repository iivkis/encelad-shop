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
	"github.com/go-playground/validator/v10"
)

type UserModelReponse struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
}

type UserHttpHandler struct {
	userService ports.UserService
	responder   *responder.Responder
	validate    *validator.Validate
}

func NewUserHttpHandler(userService ports.UserService) *UserHttpHandler {
	return &UserHttpHandler{
		userService: userService,
		responder:   responder.NewResponder(),
		validate:    validator.New(),
	}
}

type UserGetByIDResponse struct {
	UserModelReponse
}

func (h *UserHttpHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.responder.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.userService.GetByID(r.Context(), id)
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
		&UserGetByIDResponse{
			UserModelReponse{
				ID:        user.GetID(),
				Firstname: user.GetFirstname(),
				Lastname:  user.GetLastname(),
				Role:      user.GetRole().String(),
			},
		},
	)
}

type UserCreateRequest struct {
	Firstname string `json:"firstname" validate:"required,min=2,max=64"`
	Lastname  string `json:"lastname" validate:"required,min=2,max=64"`
	Password  string `json:"password" validate:"required,min=8,max=64"`
}

type UserCreateResponse struct {
	UserModelReponse
}

func (h *UserHttpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		h.responder.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	if err := h.validate.Struct(body); err != nil {
		h.responder.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.userService.Create(r.Context(),
		body.Firstname,
		body.Lastname,
		body.Password,
	)

	if err != nil {
		h.responder.JsonErr(w, http.StatusInternalServerError, nil)
		return
	}

	h.responder.JsonOk(
		w,
		http.StatusCreated,
		&UserCreateResponse{
			UserModelReponse{
				ID:        user.GetID(),
				Firstname: user.GetFirstname(),
				Lastname:  user.GetLastname(),
				Role:      user.GetRole().String(),
			},
		},
	)
}

type UserUpdateRequest struct {
	Firstname string `json:"firstname" validate:"required,min=2,max=64"`
	Lastname  string `json:"lastname" validate:"required,min=2,max=64"`
}

type UserUpdateResponse struct {
	UserModelReponse
}

func (h *UserHttpHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	if err := h.validate.Struct(body); err != nil {
		h.responder.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.userService.Update(
		r.Context(),
		id,
		body.Firstname,
		body.Lastname,
	)
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
			UserModelReponse{
				ID:        user.GetID(),
				Firstname: user.GetFirstname(),
				Lastname:  user.GetLastname(),
				Role:      user.GetRole().String(),
			},
		},
	)
}
