package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const RestCountriesBaseUrl = "https://restcountries.com/v3.1"

type Name struct {
	Common string `json:"common"`
}

type Country struct {
	Name    Name     `json:"name"`
	Capital []string `json:"capital"`
}

func retrieveCountriesByRegion(region string, ch chan []Country, errCh chan error) {
	response, err := http.Get(fmt.Sprintf("%s/region/%s", RestCountriesBaseUrl, region))
	if err != nil {
		log.Println("Error in http get: ", err.Error())
		ch <- nil
		errCh <- err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Println("The statusCode is not OK")
		ch <- nil
		errCh <- err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Can not read the body from response: ", err.Error())
		ch <- nil
		errCh <- err
	}

	var countries []Country
	err = json.Unmarshal(data, &countries)
	if err != nil {
		log.Println("Can not unmarshal the response (region)")
		ch <- nil
		errCh <- err
	}

	ch <- countries
	errCh <- nil
}

func retrieveCountriesByCurrency(currency string, ch chan []Country, errCh chan error) {
	response, err := http.Get(fmt.Sprintf("%s/currency/%s", RestCountriesBaseUrl, currency))
	if err != nil {
		log.Println("Error in http get: ", err.Error())
		ch <- nil
		errCh <- err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Println("The statusCode is not OK")
		ch <- nil
		errCh <- err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Can not read the body from response: ", err.Error())
		ch <- nil
		errCh <- err
	}

	var countries []Country
	err = json.Unmarshal(data, &countries)
	if err != nil {
		log.Println("Can not unmarshal the response (currency)")
		ch <- nil
		errCh <- err
	}

	ch <- countries
	errCh <- nil
}

func main() {
	// Scan the user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := scanner.Text()
	inputsSlice := strings.Split(inputs, " ")

	region := inputsSlice[0]
	currency := inputsSlice[1]

	// Call the Region API
	regionCh := make(chan []Country)
	regionErrCh := make(chan error)
	go retrieveCountriesByRegion(region, regionCh, regionErrCh)

	// Call the Currency API
	currencyCh := make(chan []Country)
	currencyErrCh := make(chan error)
	go retrieveCountriesByCurrency(currency, currencyCh, currencyErrCh)

	// Wait for the response
	regionCountries, regionErr := <-regionCh, <-regionErrCh
	currencyCountries, currencyErr := <-currencyCh, <-currencyErrCh
	// Close the unnecessary channels
	close(regionCh)
	close(regionErrCh)
	close(currencyCh)
	close(currencyErrCh)

	// Error handling for the API call
	if regionErr != nil {
		log.Fatalln("error in retrieving the region countries: ", regionErr.Error())
	}
	if currencyErr != nil {
		log.Fatalln("error in retrieving the currency countries: ", currencyErr.Error())
	}

	// Calculate the answer of the problem (find the common countries between two slices)
	var result []string
	for _, rc := range regionCountries {
		for _, cc := range currencyCountries {
			if rc.Name.Common == cc.Name.Common {
				result = append(result, rc.Capital[0])
			}
		}
	}

	fmt.Printf("%+v", result)
}
