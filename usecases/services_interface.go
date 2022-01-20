package usecases

import "assessment/domain"

type CarUseCase interface {
	Register(domain.Car) error
	ViewDetails(string) (domain.Car, error)
	GetCarsByColor(string) ([]domain.Car, error)
}
