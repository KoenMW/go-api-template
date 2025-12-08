package rest

import (
	"context"
	"go-api/adaptors/db"
	"go-api/domain/core"
	"go-api/domain/model"
	"go-api/domain/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserHandler struct {
	UserService service.UserService
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

	user, err := h.UserService.GetUserByID(id)

	if err != nil {
		http.Error(w, core.UserNotFound, http.StatusNotFound)
		return
	}

	WriteJSONResponse(w, http.StatusOK, &user)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	page, perPage := parsePagination(r)

	users, err := h.UserService.ListUsers(perPage, page)
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
	user, err := h.UserService.UpdateUser(&model.UserDTO{
		ID:    id,
		Name:  updateUserDTO.Name,
		Email: updateUserDTO.Email,
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
	user, err := h.UserService.CreateUser(&model.CreateUserDTO{
		Name:  createUserDTO.Name,
		Email: createUserDTO.Email,
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
	id, err = h.UserService.DeleteUser(id)
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
		UserService: *service.NewUserService(db.NewUserRepository(bun, ctx)),
	}
	return userHandler
}
