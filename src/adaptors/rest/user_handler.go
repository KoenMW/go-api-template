package rest

import (
	"context"
	"go-api/adaptors/db"
	"go-api/domain/model"
	"go-api/ports/repository"
	"net/http"
	"strconv"

	"github.com/uptrace/bun"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

var userHandler *UserHandler

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/users/"):]
	if idStr == "" {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.UserRepository.GetUserByID(int64(id))

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	WriteJSONResponse(w, http.StatusOK, &user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	createUserDTO, err := BodyReader[model.CreateUserDTO](r)
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, &model.ErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	user, err := h.UserRepository.CreateUser(createUserDTO)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, &model.ErrorResponse{
			Message: "Failed to create user",
			Error:   err.Error(),
		})
		return
	}

	WriteJSONResponse(w, http.StatusCreated, &user)
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
