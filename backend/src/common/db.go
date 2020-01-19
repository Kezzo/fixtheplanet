package common

import (
	"database/sql"
	"log"
	"os"
	"sort"
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

// ActiveTableRow ..
type ActiveTableRow struct {
	TableID      int
	TableName    string
	CreationTime time.Time
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
	query := "CREATE TABLE " + *tableName + "(issue_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL, repo VARCHAR(255) NOT NULL, issueNr INT NOT NULL, language VARCHAR(30) NULL) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"

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

// DeleteOldIssuesTables ..
func (db *Database) DeleteOldIssuesTables() error {
	stmt, err := db.SQLDB.Prepare("SELECT * FROM active_tables ORDER BY creation_time")

	if err != nil {
		return err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return err
	}

	defer rows.Close()

	activeTables := make([]ActiveTableRow, 0)

	for rows.Next() {
		activeTable := ActiveTableRow{}
		var timeString string

		err = rows.Scan(&activeTable.TableID, &activeTable.TableName, &timeString)

		if err != nil {
			return err
		}

		activeTable.CreationTime, err = time.Parse("2006-01-02 15:04:05", timeString)

		if err != nil {
			log.Println(err.Error())
		}

		activeTables = append(activeTables, activeTable)
	}

	err = rows.Err()

	if err != nil {
		return err
	}

	sort.SliceIsSorted(activeTables, func(i, j int) bool {
		return activeTables[i].CreationTime.Before(activeTables[j].CreationTime)
	})

	for i := 0; i < len(activeTables)-2; i++ {

		err := db.DeleteTable(activeTables[i].TableName)

		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteTable ..
func (db *Database) DeleteTable(tableName string) error {
	log.Println("Deleting table: " + tableName)
	stmt, err := db.SQLDB.Prepare("DROP TABLE IF EXISTS " + tableName)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	return db.DeleteActiveTableEntry(tableName)
}

// DeleteActiveTableEntry ..
func (db *Database) DeleteActiveTableEntry(tableName string) error {
	stmt, err := db.SQLDB.Prepare("DELETE FROM active_tables WHERE table_name = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(tableName)

	return err
}
