package inmemoryteststore

import (
	"assessment/domain"
	"errors"
	"strconv"
)

type Store struct {
	db []domain.Car
}

/*NewMemoryStore : Returns A Reference To Crude Database Which Stores
Its Data Inside The System Memory In A Slice

While It's Not Advised For Use In A Production Environment, It's Good
For Running Tests And Saves The Hassle Of Creating Database Mocks
*/
func NewMemoryStore() *Store {
	return &Store{db: make([]domain.Car, 0)}
}

/*Store : Saves the Car Entity Into The Database
params:
	- domain.Car
returns:
	- error
*/
func (store *Store) Store(car domain.Car) error {
	//log.Println(store.db)
	store.db = append(store.db, car)
	return nil
}

/*GetCarsByColor : Fetches The Cars Whose Colors Match The Provided Criterion
params:
	- color <string>
returns:
	- []domain.Cars
	- error
*/
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

/*GetCarByID : Fetches The Car Entity Whose ID Matches The Provided Criterion
params:
	- id <string>
returns:
	- domain.Car
	- error
*/
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

/*GetCarsByType : Fetches The Cars Whose Types Match The Provided Criterion
params:
	- carType <string>
returns:
	- []domain.Cars
	- error
*/
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
