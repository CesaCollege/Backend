package handlers

import (
	"bookman/authenticate"
	"bookman/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type signupRequestBody struct {
	Username    string    `json:"username" binding:"required"`
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	Birthday    time.Time `json:"birthday" binding:"required"`
	Password    string    `json:"password" binding:"required"`
}

type loginRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (bm *BookManagerServer) HandleLogin(c *gin.Context) {
	var lrb loginRequestBody
	err := c.ShouldBindJSON(&lrb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can not unmarshal the login request",
			"error":   err.Error(),
		})
		return
	}

	// Use authenticate package to validate the credentials
	token, err := bm.Authenticate.Login(authenticate.Credentials{
		Username: lrb.Username,
		Password: lrb.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "can not authorize the user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token.TokenString,
	})
}

func (bm *BookManagerServer) HandleSignup(c *gin.Context) {
	var srb signupRequestBody
	err := c.ShouldBindJSON(&srb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can not unmarshal the login request",
			"error":   err.Error(),
		})
		return
	}

	// Add the user to the database
	err = bm.DB.CreateNewUser(&db.User{
		Username:    srb.Username,
		FirstName:   srb.FirstName,
		LastName:    srb.LastName,
		PhoneNumber: srb.PhoneNumber,
		Birthday:    srb.Birthday,
		Password:    srb.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can not create the new user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user has been created successfully",
	})
}
