package main

import (
	"fmt"
	"golang-postgres-nginx/driver"
	"golang-postgres-nginx/handlers"
	"log"
	"net/http"
)

var handler *handlers.Repository

func main() {
	db, err := driver.NewConnection("host= port= dbname= user= password=")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to database succesfully")
	handler = handlers.NewHandler(db)
	server := http.Server{
		Addr:    ":8080",
		Handler: GetRoutes(),
	}
	fmt.Println("listens on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
