package inmemoryteststore

import (
	"assessment/domain"
	"errors"
	"strconv"
)

type Store struct {
	db []domain.Car
}

func NewMemoryStore() *Store {
	return &Store{db: make([]domain.Car, 0)}
}

func (store *Store) Store(car domain.Car) error {
	//log.Println(store.db)
	store.db = append(store.db, car)
	return nil
}

func (store *Store) GetCarsByColor(color string) ([]domain.Car, error) {
	var result []domain.Car

	for _, elem := range store.db {
		if elem.Color == color {
			result = append(result, elem)
		}
	}

	if len(result) == 0 {
		return []domain.Car{}, errors.New("no Result Found")
	}

	return result, nil
}

func (store *Store) GetCarByID(id string) (domain.Car, error) {
	numId, err := strconv.Atoi(id)

	if err != nil || numId < 1 {
		return domain.Car{}, errors.New("invalid ID Specified")
	}

	storeLength := len(store.db)

	if numId > storeLength {
		return domain.Car{}, errors.New("no Results Found")
	}

	return store.db[numId-1], nil
}

func (store *Store) GetCarsByType(carType string) ([]domain.Car, error) {
	var result []domain.Car

	for _, elem := range store.db {
		if elem.Type == carType {
			result = append(result, elem)
		}
	}

	if len(result) == 0 {
		return []domain.Car{}, errors.New("no Result Found")
	}

	return result, nil
}
