package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func NewDb() *sql.DB {
	env, err := godotenv.Read(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	var (
		host = env["HOST"]
		port = env["PORT"]
		user = env["USER"]
		pass = env["PASSWORD"]
		db   = env["DBNAME"]
	)

	postgresql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, db)

	dbpostgres, errdb := sql.Open("postgres", postgresql)

	if errdb != nil {
		panic("not connected")
	}

	errping := dbpostgres.Ping()

	if errping != nil {
		panic("not connected")
	}

	dbpostgres.SetMaxIdleConns(5)
	dbpostgres.SetMaxOpenConns(20)
	dbpostgres.SetConnMaxLifetime(60 * time.Minute)
	dbpostgres.SetConnMaxIdleTime(10 * time.Minute)

	return dbpostgres
}
