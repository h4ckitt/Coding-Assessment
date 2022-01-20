package repository

import "assessment/domain"

type CarRepository interface {
	Store(car domain.Car) error
	GetCarsByColor(color string) ([]domain.Car, error)
	GetCarByID(id string) (domain.Car, error)
}
