package usecases

import (
	"assessment/domain"
)

type Service struct {
	SchemaRepository domain.SchemaRepository
	Logger           Logger
}

func (s *Service) Register(schema domain.Schema) (int, error) {
	if err := schema.Check(); err != nil {
		return err
	}

	id, err := s.SchemaRepository.Store(schema)

	return
}

func (s *Service) ViewDetails(id int) (domain.Schema, error) {
	car, err := s.SchemaRepository.GetById(id)

	return
}

func (s *Service) GetCarsByColor(color string) ([]domain.Schema, error) {
	cars, err := s.SchemaRepository.GetCarsByColor(color)

	return
}
