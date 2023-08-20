package main

import (
	"bookman/authenticate"
	"bookman/config"
	"bookman/db"
	"bookman/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	// Read the configuration
	var cfg config.Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err.Error())
	}

	// Create a new instance of logger
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)

	// Create a new instance of database
	gormDB, err := db.NewGormDB(cfg)
	if err != nil {
		logger.WithError(err).Fatalln("error in connecting to the database")
	}
	logger.Infoln("connected to the book manager database")

	// Create schemas of models
	err = gormDB.CreateSchema()
	if err != nil {
		logger.WithError(err).Fatalln("can not auto migrate")
	}
	logger.Infoln("migrate tables successfully")

	// Create a Redis redisClient
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
	})

	// Create a new instance of authenticate
	auth, err := authenticate.NewAuth(gormDB, 10*time.Minute, logger)
	if err != nil {
		logger.WithError(err).Fatalln("can not create an instance of authenticate")
	}

	// Create a new instance of book manager http server
	bookManagerServer := handlers.BookManagerServer{
		DB:           gormDB,
		Logger:       logger,
		Authenticate: auth,
		RedisClient:  redisClient,
	}

	engine := gin.Default()

	engine.Use(auth.GinMiddleware())

	engine.POST("/auth/signup", bookManagerServer.HandleSignup)
	engine.POST("/auth/login", bookManagerServer.HandleLogin)
	engine.GET("/profile", bookManagerServer.HandleProfile)
	engine.GET("/profile/timeout", bookManagerServer.HandleProfileWithTimeout)
	engine.GET("/profile/calculate_age", bookManagerServer.HandleCalculateAge)

	engine.Run(":8080")

	logger.WithError(http.ListenAndServe(":8080", nil)).Fatalln("can not run the http server")
}
