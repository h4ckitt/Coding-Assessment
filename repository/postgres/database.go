package database

import (
	"assessment/domain"
	"assessment/usecases"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type SchemaRepository struct {
	conn *sql.DB
}

func NewPostgresHandler() (usecases.SchemaRepository, error) {
	fmt.Println("Here")
	name := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASS")
	port := os.Getenv("DATABASE_PORT")
	host := os.Getenv("DATABASE_HOST")

	db_port, _ := strconv.Atoi(port)

	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=disable", host, db_port, user, password, name)

	db, err := sql.Open("postgres", connectionInfo)

	if err != nil {
		fmt.Println("Here")
		return &SchemaRepository{}, err
	}

	err = db.Ping()

	if err != nil {
		return &SchemaRepository{}, err
	}

	return &SchemaRepository{db}, nil
}

func (handler *SchemaRepository) Store(car domain.Schema) error {
	sqlStatement := `INSERT INTO cars (type, name, color, speed_range, created_time, last_updated) VALUES ($1, $2, $3, $4, $5, $6);`

	_, err := handler.conn.Query(sqlStatement, car.Type, car.Name, car.Color, car.SpeedRange, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (handler *SchemaRepository) GetCarsByColor(color string) ([]domain.Schema, error) {
	fmt.Println("Alive")
	var cars []domain.Schema
	sqlStatement := `SELECT name, type, color, speed_range FROM cars WHERE color = ($1);`

	res, err := handler.conn.Query(sqlStatement, color)

	if err != nil {
		return []domain.Schema{}, err
	}

	defer res.Close()

	for res.Next() {
		var car domain.Schema

		res.Scan(&car.Name, &car.Type, &car.Color, &car.SpeedRange)
		fmt.Println("Appending")

		cars = append(cars, car)
	}
	fmt.Println("Returing")

	return cars, nil
}

func (handler *SchemaRepository) GetCarByID(id string) (domain.Schema, error) {
	var car domain.Schema
	sqlStatement := `SELECT name, type, color, speed_range FROM cars WHERE id = ($1);`

	row := handler.conn.QueryRow(sqlStatement, id)

	if err := row.Scan(&car.Name, &car.Type, &car.Color, &car.SpeedRange); err != nil {
		return domain.Schema{}, err
	}

	return car, nil
}
