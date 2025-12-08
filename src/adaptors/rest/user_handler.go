package rest

import (
	"context"
	"go-api/adaptors/db"
	"go-api/domain/core"
	"go-api/domain/model"
	"go-api/ports/repository"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

var userHandler *UserHandler

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, core.MissingId, http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		http.Error(w, core.InvalidId, http.StatusBadRequest)
		return
	}

	user, err := h.UserRepository.GetByID(id)

	if err != nil {
		http.Error(w, core.UserNotFound, http.StatusNotFound)
		return
	}

	WriteJSONResponse(w, http.StatusOK, &user)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	page, perPage := parsePagination(r)

	users, err := h.UserRepository.List(perPage, page)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, &model.ErrorResponse{
			Message: "Failed to list users",
			Error:   err.Error(),
		})
		return
	}

	WriteJSONResponse(w, http.StatusOK, &users)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, core.MissingId, http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, core.InvalidId, http.StatusBadRequest)
		return
	}
	updateUserDTO, err := BodyReader[model.UserDTO](r)
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, &model.ErrorResponse{
			Message: core.InvalidRequestBody,
			Error:   err.Error(),
		})
		return
	}
	user, err := h.UserRepository.Update(&model.User{
		ID:        id,
		Name:      updateUserDTO.Name,
		Email:     updateUserDTO.Email,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, &model.ErrorResponse{
			Message: "Failed to update user",
			Error:   err.Error(),
		})
		return
	}

	WriteJSONResponse(w, http.StatusOK, &user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	createUserDTO, err := BodyReader[model.CreateUserDTO](r)
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, &model.ErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	user, err := h.UserRepository.Create(&model.User{
		ID:        uuid.New(),
		Name:      createUserDTO.Name,
		Email:     createUserDTO.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, &model.ErrorResponse{
			Message: "Failed to create user",
			Error:   err.Error(),
		})
		return
	}

	WriteJSONResponse(w, http.StatusCreated, &user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, core.MissingId, http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, core.InvalidId, http.StatusBadRequest)
		return
	}
	id, err = h.UserRepository.Delete(id)
	if err != nil {
		http.Error(w, core.UserNotFound, http.StatusNotFound)
		return
	}

	WriteJSONResponse(w, http.StatusOK, &map[string]string{"id": id.String()})
}

func NewUserHandler(bun *bun.DB) *UserHandler {
	if userHandler != nil {
		return userHandler
	}

	ctx := context.Background()

	userHandler = &UserHandler{
		UserRepository: db.NewUserRepository(bun, ctx),
	}
	return userHandler
}
