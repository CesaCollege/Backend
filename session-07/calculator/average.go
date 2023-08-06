package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type averageRequest struct {
	Numbers []float64 `json:"numbers"`
}

type averageResponse struct {
	Answer float64 `json:"answer"`
}

func HandleAverage(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	reqData, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	// Unmarshalling the request body
	var body averageRequest
	err = json.Unmarshal(reqData, &body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	// Calculate the average of the number
	sum := 0.0
	for _, n := range body.Numbers {
		sum += n
	}
	answer := sum / float64(len(body.Numbers))

	// Create an instance of averageResponse
	res := averageResponse{Answer: answer}
	resData, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Write the response data
	w.WriteHeader(http.StatusOK)
	w.Write(resData)
}
