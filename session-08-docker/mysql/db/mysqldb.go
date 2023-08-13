package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-mysql/config"
)

func Connect(dbCFG *config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbCFG.DB.User, dbCFG.DB.Password, dbCFG.DB.Host, dbCFG.DB.Port, dbCFG.DB.DbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
