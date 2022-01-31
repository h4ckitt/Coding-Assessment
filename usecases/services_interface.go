package usecases

import (
	"assessment/domain"
	"net/http"
)

type CarUseCase interface {
	Register(domain.Car) error
	ViewDetails(string) (domain.Car, error)
	GetCarsByColor(string) ([]domain.Car, error)
	GetCarsByType(string) ([]domain.Car, error)
}

type Logger interface {
	LogError(string, ...interface{})
	LogAccess(*http.Request, int)
}
