package common

import (
	"database/sql"
	"log"
	"os"

	// sql driver
	_ "github.com/go-sql-driver/mysql"
)

// Database ..
type Database struct {
	SQLDB *sql.DB
}

// InitDatabase ..
func InitDatabase() (*Database, error) {
	db, err := sql.Open("mysql", os.Getenv("DB-DSN"))

	if err != nil {
		return nil, err
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	log.Println("Connection to database verified!")

	return &Database{
		SQLDB: db,
	}, nil
}
