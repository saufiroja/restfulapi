package app

import (
	"database/sql"
	"fmt"
	"restapi-golang/helper"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "kadatahu"
	dbname   = "restful-api-go"
)

func NewDB() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
