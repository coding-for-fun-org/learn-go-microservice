package main

import (
	"auth/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	conn := connectToDB()
	if conn == nil {
		log.Panic("Could not connect to Postgres")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	// define http Server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Println("Before Server started")
	// start the Server
	err := srv.ListenAndServe()
	log.Println("Server started")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Auth server running on port %s\n", webPort)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err == nil {
			log.Println("Connected to Postgres")

			return connection
		}

		log.Println("Postgres not ready yet")
		counts++

		if counts > 10 {
			log.Println("Postgres not ready after 10 attempts")

			return nil
		}

		log.Println("Backing off for 2 seconds")
		time.Sleep(2 * time.Second)

		continue
	}
}
