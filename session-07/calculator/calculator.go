package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const (
	OperatorMultiply = "*"
	OperatorDivision = "/"
	OperatorAddition = "+"
	OperatorMinus    = "-"
)

type calculationResponse struct {
	Answer int `json:"answer"`
}

func HandleCalculation(w http.ResponseWriter, r *http.Request) {
	// Parse a and b query params
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	if a == "" || b == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("a or b query params were not available")
		return
	}

	// Convert a and b to integer values
	aInt, err := strconv.Atoi(a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("a query param can not be converted to integer")
		return
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("b query param can not be converted to integer")
		return
	}

	// Get the operator from header
	op := r.Header.Get("Operator")
	var answer int
	switch op {
	case OperatorMultiply:
		answer = aInt * bInt
		break
	case OperatorDivision:
		if bInt == 0 {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("can not divide a number by zero")
			return
		}
		answer = aInt / bInt
		break
	case OperatorAddition:
		answer = aInt + bInt
		break
	case OperatorMinus:
		answer = aInt - bInt
		break
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("the operator is not valid")
		return
	}

	// Create the result response
	result := &calculationResponse{Answer: answer}
	data, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the response body
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
