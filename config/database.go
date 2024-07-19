package config

import (
	"database/sql"
	"fmt"
	"golang_listrik/helper"
	"time"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func NewDb() *sql.DB {
	env, err := godotenv.Read(".env")

	helper.Err(err)

	var (
		host = env["HOST"]
		port = env["PORT"]
		user = env["USER"]
		pass = env["PASSWORD"]
		db   = env["DBNAME"]
	)

	fmt.Println(pass)

	postgresql := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", user, pass, host, port, db)

	dbpostgres, errdb := sql.Open("postgres", postgresql)

	helper.Err(errdb)
	errping := dbpostgres.Ping()

	fmt.Println(errping)

	dbpostgres.SetMaxIdleConns(5)
	dbpostgres.SetMaxOpenConns(20)
	dbpostgres.SetConnMaxLifetime(60 * time.Minute)
	dbpostgres.SetConnMaxIdleTime(10 * time.Minute)

	return dbpostgres
}
