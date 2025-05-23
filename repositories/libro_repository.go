package repositories

import (
	"context"

	"github.com/rflorezeam/libro-read/config"
	"github.com/rflorezeam/libro-read/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroRepository interface {
	ObtenerLibros() ([]models.Libro, error)
}

type libroRepository struct {
	collection *mongo.Collection
}

func NewLibroRepository() LibroRepository {
	return &libroRepository{
		collection: config.GetCollection(),
	}
}

func (r *libroRepository) ObtenerLibros() ([]models.Libro, error) {
	var libros []models.Libro
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &libros); err != nil {
		return nil, err
	}

	return libros, nil
} 