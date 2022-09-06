package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func initDB() *sql.DB {
	conn := connectDB()
	if conn == nil {
		log.Panic("Cannot Connect to Database")
	}
	return conn
}

func connectDB() *sql.DB {
	counts := 0
	for {
		conn, err := openDB()
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Cconnection Successful")
			return conn
		}

		if counts > 10 {
			return nil
		}
		log.Println("sleep.. 1sec")
		time.Sleep(1 * time.Second)
		counts++
	}
}

func openDB() (*sql.DB, error) {
	//will be deleted once MakeFile is configured
	os.Setenv("POSTGRES", "host=localhost port=5432 user=postgres password=password dbname=e_commerce sslmode=disable timezone=UTC connect_timeout=5")

	db, err := sql.Open("postgres", os.Getenv("POSTGRES"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
