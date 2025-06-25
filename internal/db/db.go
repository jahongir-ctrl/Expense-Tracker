package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sqlx.DB

func ConnectDB() error {
	dsn := "host=localhost port=5432 user=admin password=jhg_3399 dbname=expense_tracker sslmode=disable"
	var err error
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to DB !")
	return nil
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func GetDBConn() *sqlx.DB {
	return db
}
