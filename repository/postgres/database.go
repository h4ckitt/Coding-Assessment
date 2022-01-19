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
	sqlStatement := `INSERT INTO cars (type, name, color, speed_range, created_time, last_updated) VALUES ($1, $2, $3, $4, $5, $6);`

	_, err := handler.conn.Query(sqlStatement, car.Type, car.Name, car.Color, car.SpeedRange, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (handler *SchemaRepository) GetCarsByColor(color string) ([]domain.Car, error) {
	var cars []domain.Car
	sqlStatement := `SELECT name, type, color, speed_range FROM cars WHERE color = ($1);`

	res, err := handler.conn.Query(sqlStatement, color)

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

		res.Scan(&car.Name, &car.Type, &car.Color, &car.SpeedRange)

		cars = append(cars, car)
	}

	return cars, nil
}

func (handler *SchemaRepository) GetCarByID(id string) (domain.Car, error) {
	var car domain.Car
	sqlStatement := `SELECT name, type, color, speed_range FROM cars WHERE id = ($1);`

	row := handler.conn.QueryRow(sqlStatement, id)

	if err := row.Scan(&car.Name, &car.Type, &car.Color, &car.SpeedRange); err != nil {
		return domain.Car{}, err
	}

	return car, nil
}
