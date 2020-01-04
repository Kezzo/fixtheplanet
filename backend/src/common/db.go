package common

import (
	"database/sql"
	"log"

	// sql driver
	_ "github.com/go-sql-driver/mysql"
)

// Database ..
type Database struct {
	SQLDB *sql.DB
}

// InitDatabase ..
func InitDatabase() (*Database, error) {
	db, err := sql.Open("mysql", "root:local-password@tcp(127.0.0.1:3306)/fixtheplanet")

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
