package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Repository struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	response := struct {
		Message string `json:"message"`
	}{
		Message: "your compose file now works fine",
	}
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
