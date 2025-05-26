package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rflorezeam/libro-read/services"
)

type Handler struct {
	service services.LibroService
}

func NewHandler(service services.LibroService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ObtenerLibros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	libros, err := h.service.ObtenerLibros()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(libros)
} 