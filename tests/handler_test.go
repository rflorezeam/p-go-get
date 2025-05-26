package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rflorezeam/libro-read/handlers"
	"github.com/rflorezeam/libro-read/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLibroService struct {
	mock.Mock
}

func (m *MockLibroService) ObtenerLibros() ([]models.Libro, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Libro), args.Error(1)
}

func TestObtenerLibros_Exitoso(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

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

	mockService.On("ObtenerLibros").Return(expectedLibros, nil)

	req := httptest.NewRequest(http.MethodGet, "/libros", nil)
	w := httptest.NewRecorder()

	// Act
	handler.ObtenerLibros(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response []models.Libro
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedLibros, response)
	mockService.AssertExpectations(t)
}

func TestObtenerLibros_Error(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	expectedError := errors.New("error al obtener libros")
	mockService.On("ObtenerLibros").Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, "/libros", nil)
	w := httptest.NewRecorder()

	// Act
	handler.ObtenerLibros(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedError.Error(), response["error"])
	mockService.AssertExpectations(t)
}

func TestObtenerLibros_ListaVacia(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	var expectedLibros []models.Libro
	mockService.On("ObtenerLibros").Return(expectedLibros, nil)

	req := httptest.NewRequest(http.MethodGet, "/libros", nil)
	w := httptest.NewRecorder()

	// Act
	handler.ObtenerLibros(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response []models.Libro
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Empty(t, response)
	mockService.AssertExpectations(t)
} 