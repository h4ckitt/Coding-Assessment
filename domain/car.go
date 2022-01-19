package domain

import (
	"errors"
	"time"
)

type Car struct {
	ID          int
	LastUpdated time.Time
	CreatedTime time.Time
	Type        string `json:"type"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	SpeedRange  int    `json:"speed_range"`
	Features    []string
}

type CarRepository interface {
	Store(car Car) error
	GetCarsByColor(color string) ([]Car, error)
	GetCarByID(id string) (Car, error)
}

type CarUseCase interface {
	Register(Car) error
	ViewDetails(string) (Car, error)
	GetCarsByColor(string) ([]Car, error)
}

func (s *Car) Check() error {
	if s.Name == "" {
		return errors.New("vehicle Name Cannot Be Empty")
	}

	if !contains(s.Type, []string{"sedan", "van", "suv", "motor-bike"}) {
		return errors.New("vehicle type can only be one of sedan, van, suv, motor-bike")
	}

	if !contains(s.Color, []string{"red", "green", "blue"}) {
		return errors.New("vehicle color can only be one of red, green or blue")
	}

	if s.SpeedRange > 240 || s.SpeedRange < 0 {
		return errors.New("vehicle speedrange should be between 0 and 240")
	}

	for _, elem := range s.Features {
		if !contains(elem, []string{"sunroof", "panorama", "auto-parking", "surround-system"}) {
			return errors.New("vehicle does not offer option " + elem)
		}
	}

	return nil
}

func contains(value string, slice []string) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}
	return false
}
