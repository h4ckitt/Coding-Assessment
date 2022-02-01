package repository

import (
	"assessment/domain"
)

//CarRepository : Interfaces Which Must Be Implemented By Any Data Repository
type CarRepository interface {
	Store(car domain.Car) error
	GetCarsByColor(color string) ([]domain.Car, error)
	GetCarByID(id string) (domain.Car, error)
	GetCarsByType(string) ([]domain.Car, error)
}
