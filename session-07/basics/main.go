package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type signupRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
}

func HandlePingAPI(w http.ResponseWriter, r *http.Request) {
	////////////////////////////////////////////////////////
	// Request
	////////////////////////////////////////////////////////

	// Header
	lang := r.Header.Get("Content-Language")
	if lang == "" {
		lang = "en"
	}
	fmt.Println("language: ", lang)

	// Method
	fmt.Printf("%+v\n", r.Method)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Query Params
	queryParams := r.URL.Query()
	fmt.Printf("Query Params, %+v\n", queryParams)

	// Body
	reqData, _ := io.ReadAll(r.Body)
	var sr signupRequest
	err := json.Unmarshal(reqData, &sr)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("%+v\n", sr)

	////////////////////////////////////////////////////////
	// Writer
	////////////////////////////////////////////////////////
	userInfo := map[string]interface{}{
		"name":   "mohammad",
		"age":    25,
		"gender": "male",
	}
	data, _ := json.Marshal(userInfo)

	w.Header().Add("Name", "Ali")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(data)
}

func HandleMirrorAPI(w http.ResponseWriter, r *http.Request) {
	// Read the signupRequest data
	reqData, _ := io.ReadAll(r.Body)
	var sr signupRequest
	err := json.Unmarshal(reqData, &sr)
	if err != nil {
		log.Println(err.Error())
	}

	// Write the exact same data
	resData, _ := json.Marshal(sr)
	w.Write(resData)
}

type testStruct struct {
}

func (t testStruct) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello API"))
}

func main() {
	http.Handle("/test", &testStruct{})
	http.HandleFunc("/ping", HandlePingAPI)
	http.HandleFunc("/mirror", HandleMirrorAPI)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
