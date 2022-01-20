package usecases

import (
	"assessment/domain"
)

type Service struct {
	CarRepository domain.CarRepository
	Logger        Logger
}

func NewService(repo domain.CarRepository, logger Logger) *Service {
	return &Service{repo, logger}
}

func (s *Service) Register(car domain.Car) error {
	if err := car.Check(); err != nil {
		return err
	}

	err := s.CarRepository.Store(car)

	return err
}

func (s *Service) ViewDetails(id string) (domain.Car, error) {
	car, err := s.CarRepository.GetCarByID(id)

	return car, err
}

func (s *Service) GetCarsByColor(color string) ([]domain.Car, error) {
	cars, err := s.CarRepository.GetCarsByColor(color)

	return cars, err
}
