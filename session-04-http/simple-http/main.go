package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type NameInfo struct {
	Common string `json:"common"`
}

type Country struct {
	Population int      `json:"population"`
	NameInfo   NameInfo `json:"name"`
}

func FetchCountriesWithNewRequest() {
	// Create an HTTP client
	client := http.DefaultClient

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, "https://restcountries.com/v3.1/all", nil)
	if err != nil {
		log.Fatalln("can not create the request")
	}

	// Perform the GET request
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("error in http get: ", err.Error())
	}
	defer res.Body.Close()
	log.Println("status Code: ", res.StatusCode)

	// Check the status code
	if res.StatusCode != http.StatusOK {
		log.Fatalln("the status code is not OK")
	}

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("can not read the body from response: ", err.Error())
	}

	// Unmarshall the data into Country struct
	var countries []map[string]interface{}
	err = json.Unmarshal(data, &countries)
	if err != nil {
		log.Fatalln("can not unmarshall the response")
	}

	sum := 0.0
	for _, c := range countries {
		sum += c["population"].(float64)
		fmt.Println(c["name"].(map[string]interface{})["common"], c["population"])
	}
}

func FetchCountries() {
	// Create a new request to get the countries
	res, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		log.Fatalln("error in http get: ", err.Error())
	}
	defer res.Body.Close()
	log.Println("status Code: ", res.StatusCode)

	// Check the status code
	if res.StatusCode != http.StatusOK {
		log.Fatalln("the status code is not OK")
	}

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("can not read the body from response: ", err.Error())
	}

	// Unmarshall the data into Country struct
	var countries []Country
	err = json.Unmarshal(data, &countries)
	if err != nil {
		log.Fatalln("can not unmarshall the response")
	}

	sum := 0
	for _, c := range countries {
		sum += c.Population
		fmt.Println(c.NameInfo.Common, c.Population)
	}
	//fmt.Println("The average population of the countries is ", float64(sum)/float64(len(countries)))
}

func main() {
	FetchCountriesWithNewRequest()
}
