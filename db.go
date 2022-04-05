package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dbName = os.Getenv("DB_NAME")
var dbUsername = os.Getenv("DB_USERNAME")
var dbPassword = os.Getenv("DB_PASSWORD")
var botUsername = os.Getenv("BOT_USERNAME")

//This function opens a connection to the database.
func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@/"+dbName)
	errorChecking(err)

	log.Println("********")
	log.Println("Connected to the database.")
	log.Println("********")

	return db
}
