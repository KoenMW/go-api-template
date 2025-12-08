package rest

import (
	"go-api/adaptors/db"
	"go-api/domain/core"
	"go-api/domain/model"
	"go-api/domain/service"
	"net/http"

	"github.com/uptrace/bun"
)

type UserHandler struct {
	UserService service.UserService
}

var userHandler *UserHandler

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := ParsePathID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.UserService.GetUserByID(r.Context(), id)

	if err != nil {
		http.Error(w, core.UserNotFound, http.StatusNotFound)
		return
	}

	WriteJSONResponse(w, http.StatusOK, &user)
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	page, perPage := parsePagination(r)

	users, err := h.UserService.ListUsers(r.Context(), perPage, page)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, &model.ErrorResponse{
			Message: "Failed to list users",
			Error:   err.Error(),
		})
		return
	}

	WriteJSONResponse(w, http.StatusOK, &users)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := ParsePathID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	user, err := h.UserService.UpdateUser(r.Context(), &model.UserDTO{
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

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	createUserDTO, err := BodyReader[model.CreateUserDTO](r)
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, &model.ErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	user, err := h.UserService.CreateUser(r.Context(), &model.CreateUserDTO{
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

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := ParsePathID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err = h.UserService.DeleteUser(r.Context(), id)
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

	userHandler = &UserHandler{
		UserService: *service.NewUserService(db.NewUserRepository(bun)),
	}
	return userHandler
}
