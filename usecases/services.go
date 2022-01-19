package usecases

import (
	"assessment/domain"
)

type Service struct {
	SchemaRepository SchemaRepository
	Logger           Logger
}

func NewService(repo SchemaRepository, logger Logger) *Service {
	return &Service{repo, logger}
}

func (s *Service) Register(schema domain.Schema) error {
	if err := schema.Check(); err != nil {
		return err
	}

	err := s.SchemaRepository.Store(schema)

	return err
}

func (s *Service) ViewDetails(id string) (domain.Schema, error) {
	car, err := s.SchemaRepository.GetCarByID(id)

	return car, err
}

func (s *Service) GetCarsByColor(color string) ([]domain.Schema, error) {
	cars, err := s.SchemaRepository.GetCarsByColor(color)

	return cars, err
}
