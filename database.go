package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func connectDb() {
    host     := os.Getenv("HOST")
    port     := os.Getenv("PORT")
    user     := os.Getenv("DB_USER")
    password := os.Getenv("PASSWORD")
    dbname   := os.Getenv("DBNAME")

    conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    var err error

    db, err = sql.Open("postgres", conn)
    if err != nil {
        log.Fatal("error at opening database: " + err.Error())
    }
}

func loadEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("could not load .env: " + err.Error())
    }
}
