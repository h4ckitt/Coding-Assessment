/*
	Although This Test Suite Might Not Seem Necessary As
	Establishing A Connection To The Db Might Be Deemed
	Already Enough, It's Probable That The DB To Be Connected To
	Doesn't Exist Or Isn't Connected. It's Easy To Overlook These Kind Of
	Mistakes But Tests Exist To Help Cover The Mistakes

*/

package database_test

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

var db *sql.DB

const createString = `create table cars (
id serial PRIMARY KEY,
created_time TIMESTAMPTZ,
last_updated TIMESTAMPTZ,
name VARCHAR NOT NULL,
type VARCHAR NOT NULL,
color VARCHAR NOT NULL,
speed_range INTEGER NOT NULL
)`

const dropString = `DROP TABLE IF EXISTS cars`

type DbTestSuite struct {
	suite.Suite
	m *migrate.Migrate
}

func (s *DbTestSuite) SetupSuite() {
	var err error

	name := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASS")
	port := os.Getenv("DATABASE_PORT")
	host := os.Getenv("DATABASE_HOST")

	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable", host, port, user, password, name)

	db, err = sql.Open("postgres", connectionInfo)

	require.NoError(s.T(), err)

}

func TestDbTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}

//Test Database Deletion
func (s *DbTestSuite) TearDownSuite() {
	_, err := db.Exec(dropString)

	require.NoError(s.T(), err)
}

//Test DB Connection By Creating A New Table
func (s *DbTestSuite) TestCreateTable() {

	_, err := db.Exec(createString)

	require.NoError(s.T(), err)
}

//Test That We're Talking To The Same DB
// By Trying To Create The Same Table Twice
// Doesn't Make Sense? (Probably)
func (s *DbTestSuite) TestCreateTableError() {
	_, err := db.Exec(createString)

	require.Error(s.T(), err)
}
