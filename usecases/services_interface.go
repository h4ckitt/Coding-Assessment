package usecases

import (
	"assessment/domain"
	"net/http"
)

//CarUseCase : Interfaces Which Must Be Implemented By Any Usecase Handler
type CarUseCase interface {
	Register(domain.Car) error
	ViewDetails(string) (domain.Car, error)
	GetCarsByColor(string) ([]domain.Car, error)
	GetCarsByType(string) ([]domain.Car, error)
}

//Logger : Interfaces That Must Be Implemented By Any Logger Service
type Logger interface {
	LogError(string, ...interface{})
	LogAccess(*http.Request, int)
}
