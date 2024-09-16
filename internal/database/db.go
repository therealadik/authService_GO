package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	connStr := "user=postgres password=postgres dbname=auth_service_db sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("d%v", err)
	}

	fmt.Println("Success connect db")
	return nil
}
