package usecases

import "domain"

type SchemaRepository interface {
	Store(car domain.Schema) (id int, err error)
	GetCarsByColor(color string) ([]domain.Schema, error)
	GetCarByID(id int) (domain.Schema, error)
}
