package common

import (
	"database/sql"
	"log"
	"os"
	"time"

	// sql driver
	_ "github.com/go-sql-driver/mysql"
)

// Database ..
type Database struct {
	SQLDB                    *sql.DB
	LastActiveTableCheckTime int64
	ActiveTable              string
}

// InitDatabase ..
func InitDatabase() (*Database, error) {
	db, err := sql.Open("mysql", os.Getenv("DB-DSN"))

	if err != nil {
		return nil, err
	}

	// Open doesn't open a connection. Validate DSN data:
	err = CheckConnection(db)

	if err != nil {
		return nil, err
	}

	log.Println("Connection to database verified!")

	return &Database{
		SQLDB: db,
	}, nil
}

// CheckConnection ..
func CheckConnection(db *sql.DB) error {
	var err error

	for try := 0; try < 5; try++ {
		log.Println("Checking connection to database")
		err = db.Ping()

		if err == nil {
			return nil
		}

		log.Println("Database not ready, trying again in 5")
		time.Sleep(5 * time.Second)
	}

	return err
}

// CheckActiveTable ..
func (db *Database) CheckActiveTable() error {
	db.LastActiveTableCheckTime = time.Now().Unix()

	stmt, err := db.SQLDB.Prepare("SELECT * FROM active_tables ORDER BY creation_time LIMIT 1")

	if err != nil {
		return err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return err
	}

	defer rows.Close()

	var tableID int
	var tableName string
	var creationTime string

	for rows.Next() {
		err = rows.Scan(&tableID, &tableName, &creationTime)

		if err != nil {
			return err
		}
	}

	err = rows.Err()

	if err != nil {
		return err
	}

	db.ActiveTable = tableName
	return nil
}

// CreateNextIssuesTable ..
func (db *Database) CreateNextIssuesTable(tableName *string) error {
	query := "CREATE TABLE " + *tableName + "(issue_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255) NOT NULL, repo VARCHAR(255) NOT NULL, issueNr INT NOT NULL, language VARCHAR(30) NULL)"

	stmt, err := db.SQLDB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	return err
}

// AddActiveTableEntry ..
func (db *Database) AddActiveTableEntry(tableName *string) error {
	stmt, err := db.SQLDB.Prepare("INSERT INTO active_tables VALUES( ?, ?, CURRENT_TIMESTAMP() )")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(nil, tableName)

	return err
}
