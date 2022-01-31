package usecases_test

import (
	"assessment/domain"
	"assessment/infrastructure/db/inmemoryteststore"
	"assessment/usecases"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceTestSuite struct {
	suite.Suite
	service usecases.CarUseCase
}

func (s *ServiceTestSuite) SetupSuite() {
	repo := inmemoryteststore.NewMemoryStore()

	cars := []domain.Car{
		{
			Name:       "Toyota Tundra",
			Color:      "red",
			Type:       "suv",
			SpeedRange: 180,
			Features:   []string{"panorama", "surround-system"},
		},
		{
			Name:       "Tesla Model S Plaid",
			Color:      "blue",
			Type:       "sedan",
			SpeedRange: 240,
			Features:   []string{"auto-parking", "panorama", "surround-system"},
		},
		{
			Name:       "Ducati 99R",
			Color:      "green",
			Type:       "motor-bike",
			SpeedRange: 200,
			Features:   []string{"surround-system"},
		},
		{
			Name:       "Ford ES500",
			Color:      "blue",
			Type:       "van",
			SpeedRange: 80,
			Features:   []string{"surround-system", "sunroof", "surround-system"},
		},
	}

	for _, car := range cars {
		err := repo.Store(car)

		require.NoError(s.T(), err)
	}

	s.service = usecases.NewService(repo)
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) TestRegister() {
	invalidTypes := []string{"bicycle", "airplane", "scooter"}
	invalidColors := []string{"white", "purple", "orange", "black"}
	invalidSpeedRanges := []int{250, -1, 800}
	invalidFeatures := []string{"parking-aid", "seat-heating", "airplay"}

	car := domain.Car{
		Name:       "Toyota Prius",
		Color:      "blue",
		Type:       "sedan",
		SpeedRange: 220,
		Features:   []string{"sunroof", "panorama", "surround-system"},
	}

	err := s.service.Register(car)

	require.NoError(s.T(), err)

	for _, types := range invalidTypes {
		car.Type = types

		err := s.service.Register(car)

		require.Errorf(s.T(), err, "Expected An Error, Got: %v\n", err)
	}

	car.Type = "sedan"

	for _, color := range invalidColors {
		car.Color = color

		err := s.service.Register(car)
		require.Errorf(s.T(), err, "Expected An Error, Got: %v\n", err)
	}

	car.Color = "blue"

	for _, speed := range invalidSpeedRanges {
		car.SpeedRange = speed

		err := s.service.Register(car)
		require.Errorf(s.T(), err, "Expected An Error, Got: %v\n", err)
	}

	car.SpeedRange = 240
	car.Features = invalidFeatures

	err = s.service.Register(car)
	require.Errorf(s.T(), err, "Expected An Error, Got: %v\n", err)

}

func (s *ServiceTestSuite) TestViewDetails() {
	_, err := s.service.ViewDetails("2")

	require.NoError(s.T(), err, "Expected: nil, Got: %v\n", err)

	for _, id := range []string{"-1", "250", "10000000000000"} {
		_, err := s.service.ViewDetails(id)
		require.Errorf(s.T(), err, "Expected An Error, Got: %v\n", err)
	}
}

func (s *ServiceTestSuite) TestGetCarsByColor() {
	_, err := s.service.GetCarsByColor("blue")

	require.NoError(s.T(), err, "Expected: nil, Got: %v\n", err)

	_, err = s.service.GetCarsByColor("white")

	require.Errorf(s.T(), err, "Expected An Error, Got: %v\n", err)
}

func (s *ServiceTestSuite) TestGetCarsByType() {
	_, err := s.service.GetCarsByType("sedan")

	require.NoError(s.T(), err, "Expected: nil, Got: %v\n", err)

	_, err = s.service.GetCarsByType("white")

	require.Errorf(s.T(), err, "Expected An Error, Got: %v\n", err)

	_, err = s.service.GetCarsByType("airplane")

	require.Errorf(s.T(), err, "Expected An Error, Got: %v\n", err)
}
