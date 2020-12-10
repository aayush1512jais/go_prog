package repository

import (
	"database/sql"
	"errors"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

var Err = errors.New("False Error")

type AppRepository interface {
	Get(id int) (int, error)
	Add(id int) error
}

func (db *Database) Get(id int) (int, error) {
	return id, Err
}
func (db *Database) Add(id int) error {
	return Err
}
