package handlers

import (
	"bookman/authenticate"
	"bookman/db"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type BookManagerServer struct {
	DB           *db.GormDB
	Logger       *logrus.Logger
	Authenticate *authenticate.Auth
	RedisClient  *redis.Client
}
