package rest

import (
	"encoding/json"
	"go-api/domain"
	"go-api/ports"
	"net/http"
)

type Handler struct {
	Producer ports.HelloWorldProducer
}

func (h *Handler) Messages(w http.ResponseWriter, r *http.Request) {
	var msg domain.Message
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
