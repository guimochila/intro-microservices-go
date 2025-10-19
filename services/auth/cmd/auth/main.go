// Copyright (c) 2025, guimochila.com. Continuous Learning.

package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPassword = "Admin123"
	dbName     = "user"
)

var db *sql.DB

func main() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPassword, dbName)

	db, err = sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing db: %s", err)
		}
	}()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
