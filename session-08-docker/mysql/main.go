package main

import (
	"golang-mysql/config"
	"golang-mysql/db"
	"log"
)

func main() {
	dbConfig := &config.Config{
		DB: config.MysqlDB{
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
