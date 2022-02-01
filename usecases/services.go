package usecases

import (
	"assessment/domain"
	"assessment/repository"
)

type Service struct {
	CarRepository repository.CarRepository
}

//NewService : Returns A New Service Entity Which Implements The Usecase Interface
func NewService(repo repository.CarRepository) *Service {
	return &Service{repo}
}

/*Register : Communicates The Register Request To The Repository And Does Pre-Flight Checks
params:
	- domain.car

returns:
	- error
*/
func (s *Service) Register(car domain.Car) error {
	if err := car.Check(); err != nil {
		return err
	}

	err := s.CarRepository.Store(car)

	return err
}

/*ViewDetails : Communicates Fetching Of The Car Whose ID Matches The Provided ID To The Rep
params:
	- id <string>

returns:
	- error
*/
func (s *Service) ViewDetails(id string) (domain.Car, error) {
	car, err := s.CarRepository.GetCarByID(id)

	return car, err
}

/*GetCarsByColor : Communicates Fetching Of Cars Whose Colors Match The Provided Color To The Repo
params:
	- color <string>
returns:
	- []domain.Car
	- error
*/
func (s *Service) GetCarsByColor(color string) ([]domain.Car, error) {
	cars, err := s.CarRepository.GetCarsByColor(color)

	return cars, err
}

/*GetCarsByType : Communicates Fetching Of Cars Whose Types Match The Provided Color To The Repo
params:
	- carType <string>
returns:
	- []domain.Car
	- error
*/
func (s *Service) GetCarsByType(carType string) ([]domain.Car, error) {
	cars, err := s.CarRepository.GetCarsByType(carType)

	return cars, err
}
