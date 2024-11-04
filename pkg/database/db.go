package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
