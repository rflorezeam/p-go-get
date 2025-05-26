package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/libro-read/config"
	"github.com/rflorezeam/libro-read/handlers"
	"github.com/rflorezeam/libro-read/repositories"
	"github.com/rflorezeam/libro-read/services"
)

func main() {
	// Inicializar la base de datos
	config.ConectarDB()

	// Inicializar las capas
	repo := repositories.NewLibroRepository()
	service := services.NewLibroService(repo)
	handler := handlers.NewHandler(service)
	
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