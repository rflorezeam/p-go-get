package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/libro-read/config"
	"github.com/rflorezeam/libro-read/repositories"
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

	json.NewEncoder(w).Encode(libros)
}

func main() {
	// Inicializar la base de datos
	config.ConectarDB()

	// Inicializar las capas
	repo := repositories.NewLibroRepository()
	service := services.NewLibroService(repo)
	handler := NewHandler(service)
	
	// Configurar el router
	router := mux.NewRouter()
	router.HandleFunc("/libros", handler.ObtenerLibros).Methods("GET")

	// Configurar el puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	fmt.Printf("Servicio de lectura de libros corriendo en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
} 