package database

import (
	"assessment/domain"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type SchemaRepository struct {
	conn *sql.DB
}

func NewPostgresHandler() (domain.CarRepository, error) {
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
		var features []string
		var id int

		res.Scan(&id, &car.Name, &car.Type, &car.Color, &car.SpeedRange)

		featureRes, err := handler.conn.Query(featureFetchStatement, id)

		if err != nil {
			return []domain.Car{}, err
		}

		for featureRes.Next() {
			var feature string
			featureRes.Scan(&feature)
			features = append(features, feature)
		}
		car.Features = append(car.Features, features...)
		cars = append(cars, car)
	}

	return cars, nil
}

func (handler *SchemaRepository) GetCarByID(id string) (domain.Car, error) {
	var (
		car     domain.Car
		feature string
	)
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
