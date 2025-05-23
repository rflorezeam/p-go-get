package services

import (
	"github.com/rflorezeam/libro-read/models"
	"github.com/rflorezeam/libro-read/repositories"
)

type LibroService interface {
	ObtenerLibros() ([]models.Libro, error)
}

type libroService struct {
	repo repositories.LibroRepository
}

func NewLibroService(repo repositories.LibroRepository) LibroService {
	return &libroService{
		repo: repo,
	}
}

func (s *libroService) ObtenerLibros() ([]models.Libro, error) {
	return s.repo.ObtenerLibros()
} 