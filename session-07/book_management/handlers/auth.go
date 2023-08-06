package handlers

import (
	"bookman/authenticate"
	"bookman/db"
	"encoding/json"
	"io"
	"net/http"
)

type signupRequestBody struct {
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type loginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (bm *BookManagerServer) HandleSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body for the new user
	reqData, err := io.ReadAll(r.Body)
	if err != nil {
		bm.Logger.WithError(err).Warn("can not read the request data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var srb signupRequestBody
	err = json.Unmarshal(reqData, &srb)
	if err != nil {
		bm.Logger.WithError(err).Warn("can not unmarshal the new user request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add the user to the database
	user := &db.User{
		Username:    srb.Username,
		FirstName:   srb.FirstName,
		LastName:    srb.LastName,
		PhoneNumber: srb.PhoneNumber,
		Password:    srb.Password,
	}
	err = bm.Db.CreateNewUser(user)
	if err != nil {
		bm.Logger.WithError(err).Warn("can not create a new user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "user has been created",
	}
	resBody, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (bm *BookManagerServer) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body for the new user
	reqData, err := io.ReadAll(r.Body)
	if err != nil {
		bm.Logger.WithError(err).Warn("can not read the request data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var lrb loginRequestBody
	err = json.Unmarshal(reqData, &lrb)
	if err != nil {
		bm.Logger.WithError(err).Warn("can not unmarshal the login request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check the authentication
	token, err := bm.Authenticate.Login(authenticate.Credentials{
		Username: lrb.Username,
		Password: lrb.Password,
	})
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Create the response
	response := map[string]interface{}{
		"access_token": token.TokenString,
	}
	resBody, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
