package rest

import (
	"encoding/json"
	service "go-api/domain/service"
	repository "go-api/ports/repository"
	"net/http"
)

type Handler struct {
	Producer repository.HelloWorldProducer
}

func (h *Handler) Messages(w http.ResponseWriter, r *http.Request) {
	var msg service.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := msg.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Producer.Log(msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("message sent"))
}
