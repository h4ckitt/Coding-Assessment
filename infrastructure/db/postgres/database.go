package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"assessment/domain"

	_ "github.com/lib/pq"
)

type SchemaRepository struct {
	conn *sql.DB
}

/*NewPostgresHandler : Returns A New Database Handler Which Would Be
Used To Communicate With The Database

	returns:
		- *schemaRepository
		- error
*/
func NewPostgresHandler() (*SchemaRepository, error) {
	fmt.Println("Here")
	name := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASS")
	port := os.Getenv("DATABASE_PORT")
	host := os.Getenv("DATABASE_HOST")

	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable", host, port, user, password, name)

	db, err := sql.Open("postgres", connectionInfo)

	if err != nil {
		return &SchemaRepository{}, err
	}

	err = db.Ping()

	if err != nil {
		return &SchemaRepository{}, err
	}

	return &SchemaRepository{db}, nil
}

/*Store : Ensures That The Provided Object Is Stored Into The Database
params:
	- domain.Car

returns:
	- error
*/
func (handler *SchemaRepository) Store(car domain.Car) error {
	var id int
	carInsertStatement := `INSERT INTO cars (type, name, color, speed_range, created_time, last_updated) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	featureInsertStatement := `INSERT INTO features (car_id, feature) VALUES ($1, $2);`

	err := handler.conn.QueryRow(carInsertStatement, car.Type, car.Name, car.Color, car.SpeedRange, time.Now(), time.Now()).Scan(&id)

	if err != nil {
		return err
	}

	for _, elem := range car.Features {
		_, err := handler.conn.Exec(featureInsertStatement, id, elem)

		if err != nil {
			return err
		}
	}

	return nil
}

/*GetCarsByColor : Retrieves The Entity Whose Color Matches The One
Provided By The Caller

	params:
		- color <string>

	returns:
		- []domain.Car
		- error
*/
func (handler *SchemaRepository) GetCarsByColor(color string) ([]domain.Car, error) {
	var cars []domain.Car
	carFetchStatement := `SELECT id, name, type, color, speed_range FROM cars WHERE color = ($1);`
	featureFetchStatement := `SELECT feature FROM features WHERE car_id = $1;`

	res, err := handler.conn.Query(carFetchStatement, color)

	if err != nil {
		return []domain.Car{}, err
	}

	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)

	for res.Next() {
		var car domain.Car
		car.Features = make([]string, 0)
		var id int

		res.Scan(&id, &car.Name, &car.Type, &car.Color, &car.SpeedRange)

		featureRes, err := handler.conn.Query(featureFetchStatement, id)

		if err != nil {
			return []domain.Car{}, err
		}

		for featureRes.Next() {
			var feature string
			featureRes.Scan(&feature)
			car.Features = append(car.Features, feature)
		}

		cars = append(cars, car)
	}

	return cars, nil
}

/*GetCarsByType : Retrieves The Entity Whose type Matches The One
Provided By The Caller

	params:
		- type <string>

	returns:
		- []domain.Car
		- error
*/
func (handler *SchemaRepository) GetCarsByType(carType string) ([]domain.Car, error) {
	var cars []domain.Car

	carFetchStatement := `SELECT id, name, type, color, speed_range FROM cars WHERE type = ($1);`
	featureFetchStatement := `SELECT feature FROM features WHERE car_id = $1;`

	res, err := handler.conn.Query(carFetchStatement, carType)

	if err != nil {
		return []domain.Car{}, err
	}

	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)

	for res.Next() {
		var car domain.Car
		car.Features = make([]string, 0)
		var id int

		res.Scan(&id, &car.Name, &car.Type, &car.Color, &car.SpeedRange)

		featureRes, err := handler.conn.Query(featureFetchStatement, id)

		if err != nil {
			return []domain.Car{}, err
		}

		for featureRes.Next() {
			var feature string
			featureRes.Scan(&feature)
			car.Features = append(car.Features, feature)
		}

		cars = append(cars, car)
	}

	return cars, nil

}

/*GetCarsByID : Retrieves The Entity Whose id Matches The One
Provided By The Caller

	params:
		- id <string>

	returns:
		- domain.Car
		- error
*/
func (handler *SchemaRepository) GetCarByID(id string) (domain.Car, error) {
	var (
		car     domain.Car
		feature string
	)
	car.Features = make([]string, 0)

	carFetchStatement := `SELECT name, type, color, speed_range FROM cars WHERE id = ($1);`
	featureFetchStatement := `SELECT feature FROM features WHERE car_id = $1`

	carRow := handler.conn.QueryRow(carFetchStatement, id)
	featureRow, err := handler.conn.Query(featureFetchStatement, id)

	if err != nil {
		return domain.Car{}, err
	}

	if err := carRow.Scan(&car.Name, &car.Type, &car.Color, &car.SpeedRange); err != nil {
		return domain.Car{}, err
	}

	for featureRow.Next() {
		featureRow.Scan(&feature)

		car.Features = append(car.Features, feature)
	}

	return car, nil
}
