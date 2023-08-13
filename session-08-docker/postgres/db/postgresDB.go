package db

import (
	"database/sql"
	"fmt"
	"golang-postgres/config"

	_ "github.com/lib/pq"
)

func Connect(dbCFG *config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbCFG.DB.Host, dbCFG.DB.Port, dbCFG.DB.User, dbCFG.DB.Password, dbCFG.DB.DbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
