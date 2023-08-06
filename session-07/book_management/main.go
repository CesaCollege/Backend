package main

import (
	"bookman/authenticate"
	"bookman/config"
	"bookman/db"
	"bookman/handlers"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	// Load the config
	var cfg config.Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err.Error())
	}

	// Setup the logger
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	//logger.SetFormatter(&logrus.JSONFormatter{})

	// Create new instance of dg.DB
	gormDB, err := db.NewGormDB(cfg)
	if err != nil {
		logger.WithError(err).Fatal("error in connecting to the postgres database")
	}
	logger.Info("connected to the mashgh database")

	err = gormDB.CreateSchemas()
	if err != nil {
		logger.WithError(err).Fatal("error in database migration")
	}
	logger.Infoln("migrate tables and models successfully")

	// Create authenticate
	auth, err := authenticate.NewAuth(gormDB, 10, logger)
	if err != nil {
		logger.WithError(err).Fatal("can not create the authenticate instance")
	}

	bookManagerServer := handlers.BookManagerServer{Db: gormDB, Logger: logger, Authenticate: auth}

	http.HandleFunc("/auth/signup", bookManagerServer.HandleSignup)
	http.HandleFunc("/auth/login", bookManagerServer.HandleLogin)
	http.HandleFunc("/profile", bookManagerServer.HandleProfile)

	logger.WithError(http.ListenAndServe(":8080", nil)).Fatalln("can not setup the server")
}
