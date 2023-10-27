// db/connection.go
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

var DBCOnn *sql.DB

func InitializeDB() error {
	e := godotenv.Load()
	if e != nil {
		fmt.Println("Error loading .env file:", e)
		return e
	}

	dbURL := os.Getenv("DB_URL")

	var err error
	DBCOnn, err = sql.Open("mysql", dbURL)
	if err != nil {
		return err
	}
	return nil
}
