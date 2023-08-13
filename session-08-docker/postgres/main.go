package main

import (
	"golang-postgres/config"
	"golang-postgres/db"
	"log"
)

func main() {
	dbConfig := &config.Config{
		DB: config.PostgresDB{
			Host:     "",
			Port:     0,
			DbName:   "",
			User:     "",
			Password: "",
		},
	}
	db, err := db.Connect(dbConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	log.Println("successfully connected")
}
