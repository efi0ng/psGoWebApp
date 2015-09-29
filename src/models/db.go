package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3","./db/gowebapp.sqlite")
	return db, err
}