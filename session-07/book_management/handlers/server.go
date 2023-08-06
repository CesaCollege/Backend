package handlers

import (
	"bookman/authenticate"
	"bookman/db"
	"github.com/sirupsen/logrus"
)

type BookManagerServer struct {
	Db           *db.GormDB
	Logger       *logrus.Logger
	Authenticate *authenticate.Auth
}
