package models

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func SetupDB(dbUrl string) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	DB = db
}