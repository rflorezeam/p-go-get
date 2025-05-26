package repositories

import (
	"errors"
	"testing"

	"github.com/rflorezeam/libro-read/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLibroRepository struct {
	mock.Mock
}

func (m *MockLibroRepository) ObtenerLibros() ([]models.Libro, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Libro), args.Error(1)
}

func TestObtenerLibros_RepositorioExitoso(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)

	expectedLibros := []models.Libro{
		{
			ID:     "1",
			Titulo: "Libro 1",
			Autor:  "Autor 1",
		},
		{
			ID:     "2",
			Titulo: "Libro 2",
			Autor:  "Autor 2",
		},
	}

	mockRepo.On("ObtenerLibros").Return(expectedLibros, nil)

	// Act
	libros, err := mockRepo.ObtenerLibros()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedLibros, libros)
	mockRepo.AssertExpectations(t)
}

func TestObtenerLibros_RepositorioError(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)

	expectedError := errors.New("error de conexión con la base de datos")
	mockRepo.On("ObtenerLibros").Return(nil, expectedError)

	// Act
	libros, err := mockRepo.ObtenerLibros()

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, libros)
	mockRepo.AssertExpectations(t)
}

func TestObtenerLibros_RepositorioListaVacia(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)

	var expectedLibros []models.Libro
	mockRepo.On("ObtenerLibros").Return(expectedLibros, nil)

	// Act
	libros, err := mockRepo.ObtenerLibros()

	// Assert
	assert.NoError(t, err)
	assert.Empty(t, libros)
	mockRepo.AssertExpectations(t)
} 