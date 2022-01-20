package usecases

import (
	"assessment/domain"
	"assessment/repository"
)

type Service struct {
	CarRepository repository.CarRepository
	Logger        Logger
}

func NewService(repo repository.CarRepository, logger Logger) *Service {
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
