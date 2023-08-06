package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/calculate", HandleCalculation)
	http.HandleFunc("/average", HandleAverage)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
