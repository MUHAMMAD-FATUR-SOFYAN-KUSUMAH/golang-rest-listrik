package main

import (
	"fmt"
	"golang_listrik/config"
)

func main() {
	// var db *sql.DB
	// db = config.NewDb()
	config.NewDb()
	// row, err := db.Query("SELECT version()")

	fmt.Println("status ok")
}
