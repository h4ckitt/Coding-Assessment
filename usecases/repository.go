package usecases

import "assessment/domain"

type SchemaRepository interface {
	Store(car domain.Schema) error
	GetCarsByColor(color string) ([]domain.Schema, error)
	GetCarByID(id string) (domain.Schema, error)
}
