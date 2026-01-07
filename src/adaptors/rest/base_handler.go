package rest

import (
	"go-api/domain/model"
	"go-api/ports/repository"
	"go-api/ports/service"
	"net/http"

	"github.com/google/uuid"
)

type Basehandler[T model.BaseEntity, DTO model.BaseDTO[T], CreateDTO model.BaseCreateDTO[T], Repository repository.BaseRepository[T], Service service.BaseService[T, DTO, CreateDTO, Repository]] struct {
	service Service
}

func (h *Basehandler[T, DTO, CreateDTO, Repository, Service]) Create(w http.ResponseWriter, r *http.Request) {
	createDTO, err := BodyReader[CreateDTO](r)
	if err != nil {
		WriteJSONResponse(
			w, http.StatusBadRequest, &map[string]string{
				"message": "Invalid request body",
				"error":   err.Error(),
			},
		)
		return
	}

	createdEntity, err := h.service.Create(r.Context(), createDTO)
	if err != nil {
		WriteJSONResponse(
			w, http.StatusInternalServerError, &map[string]string{
				"message": "Failed to create entity",
				"error":   err.Error(),
			},
		)
		return
	}

	WriteJSONResponse(w, http.StatusCreated, &createdEntity)
}

func (h *Basehandler[T, DTO, CreateDTO, Repository, Service]) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := ParsePathID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entity, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Entity not found", http.StatusNotFound)
		return
	}

	WriteJSONResponse(w, http.StatusOK, &entity)
}

func (h *Basehandler[T, DTO, CreateDTO, Repository, Service]) List(w http.ResponseWriter, r *http.Request) {
	page, perPage := parsePagination(r)
	entities, err := h.service.List(r.Context(), perPage, page)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, &map[string]string{
			"message": "Failed to list entities",
			"error":   err.Error(),
		})
		return
	}

	WriteJSONResponse(w, http.StatusOK, &entities)
}

func (h *Basehandler[T, DTO, CreateDTO, Repository, Service]) Update(w http.ResponseWriter, r *http.Request) {
	id, err := ParsePathID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateDTO, err := BodyReader[DTO](r)
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, &map[string]string{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	updatedEntity, err := h.service.Update(r.Context(), id, updateDTO)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, &map[string]string{
			"message": "Failed to update entity",
			"error":   err.Error(),
		})

		return
	}

	WriteJSONResponse(w, http.StatusOK, &updatedEntity)
}

func (h *Basehandler[T, DTO, CreateDTO, Repository, Service]) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := ParsePathID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deletedID, err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to delete entity", http.StatusInternalServerError)
		return
	}

	WriteJSONResponse(w, http.StatusOK, &map[string]uuid.UUID{
		"deleted_id": deletedID,
	})
}
