package repository_test

import (
	"assessment/domain"
	db "assessment/infrastructure/db/inmemoryteststore"
	"assessment/repository"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RepoTestSuite struct {
	suite.Suite
	repo repository.CarRepository
}

func (s *RepoTestSuite) SetupSuite() {
	s.repo = db.NewMemoryStore()
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}

func (s *RepoTestSuite) TestAStoreNewCar() {
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
		err := s.repo.Store(car)

		require.NoError(s.T(), err)
	}
}

func (s *RepoTestSuite) TestGetCarByID() {
	carr := domain.Car{
		Name:       "Tesla Model S Plaid",
		Color:      "blue",
		Type:       "sedan",
		SpeedRange: 240,
		Features:   []string{"auto-parking", "panorama", "surround-system"},
	}
	car, err := s.repo.GetCarByID("2")

	require.NoError(s.T(), err)
	require.Equal(s.T(), car, carr)

	ids := []string{"20", "-1", "0"}

	for _, id := range ids {
		_, err := s.repo.GetCarByID(id)
		require.Error(s.T(), err, "Expected Error, Got: %v\n", err)
	}
}

func (s *RepoTestSuite) TestGetCarsByColor() {
	_, err := s.repo.GetCarsByColor("blue")

	require.NoError(s.T(), err)

	_, err = s.repo.GetCarsByColor("white")

	require.Errorf(s.T(), err, "Expected An Error, Got: %v\n", err)
}

func (s *RepoTestSuite) TestGetCarsByType() {
	types := []string{"sedan", "van", "motor-bike", "suv"}

	for _, typee := range types {
		_, err := s.repo.GetCarsByType(typee)
		require.NoError(s.T(), err)
	}

	_, err := s.repo.GetCarsByType("bicycle")

	require.Errorf(s.T(), err, "Expected An Error, Got: %v\n", err)

}
