package handlers

import (
	"bookman/db"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"time"
)

type userInfoResponse struct {
	Username    string    `json:"username"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Birthday    time.Time `json:"birthday"`
}

func (bm *BookManagerServer) HandleProfile(c *gin.Context) {
	//// Grab authorization header
	//token := c.GetHeader("Authorization")
	//if token == "" {
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"message": "you should provide the authentication token",
	//	})
	//	return
	//}
	//
	//// Retrieve the related account by token
	//account, err := bm.Authenticate.GetAccountByToken(token)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"message": "the user token is not valid or has been expired",
	//	})
	//	return
	//}
	account, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "you should provide the authentication token",
		})
		return
	}

	// Retrieve user from database
	user, err := bm.DB.GetUserByUsername(*account.(*string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can not retrieve this user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &userInfoResponse{
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Birthday:    user.Birthday,
	})
}

func (bm *BookManagerServer) HandleCalculateAge(c *gin.Context) {
	// Grab authorization header
	//token := c.GetHeader("Authorization")
	//if token == "" {
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"message": "you should provide the authentication token",
	//	})
	//	return
	//}
	//
	//// Retrieve the related account by token
	//account, err := bm.Authenticate.GetAccountByToken(token)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"message": "the user token is not valid or has been expired",
	//	})
	//	return
	//}
	account, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "you should provide the authentication token",
		})
		return
	}

	// Search in redis for caching result
	useCache := c.GetHeader("Use-Cache")
	if useCache == "true" {
		age, err := bm.getCacheCalculatedAge(c, *account.(*string))
		if err != nil {
			bm.Logger.WithError(err).Infoln("can not use cache")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"years": age,
			})
			return
		}
	}

	// Retrieve user from database
	user, err := bm.DB.GetUserByUsername(*account.(*string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can not retrieve this user",
			"error":   err.Error(),
		})
		return
	}

	// Calculate the age
	years := math.Round((time.Now().Sub(user.Birthday).Hours() / 24) / 365)
	time.Sleep(time.Millisecond)
	if useCache == "true" {
		bm.cacheCalculatedAge(c, *account.(*string), years)
	}

	c.JSON(http.StatusOK, gin.H{
		"years": years,
	})
}

func (bm *BookManagerServer) getCacheCalculatedAge(ctx context.Context, username string) (string, error) {
	cacheKey := fmt.Sprintf("user_age:%s", username)
	return bm.RedisClient.Get(ctx, cacheKey).Result()
}

func (bm *BookManagerServer) cacheCalculatedAge(ctx context.Context, username string, years float64) {
	cacheKey := fmt.Sprintf("user_age:%s", username)
	err := bm.RedisClient.Set(ctx, cacheKey, years, time.Minute).Err()
	if err != nil {
		bm.Logger.WithError(err).Warn("can not cache the response")
	}
}

func (bm *BookManagerServer) HandleProfileWithTimeout(c *gin.Context) {
	// Set a context timeout of 5 seconds
	ctx, cancel := context.WithTimeout(c, 5*time.Microsecond)
	defer cancel() // Ensure the context is canceled to release resources

	account, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "you should provide the authentication token",
		})
		return
	}

	// Create a channel to receive the result from the database operation
	resultCh := make(chan *db.User, 1)

	// Start the database operation in a goroutine
	go func() {
		// Retrieve user from database (replace with actual query)
		user, err := bm.DB.GetUserByUsername(*account.(*string))
		if err != nil {
			resultCh <- nil
			return
		}
		resultCh <- user
	}()

	// Wait for the database operation result or the context timeout
	select {
	case <-ctx.Done():
		// Context timeout or cancellation occurred
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"message": "Request timed out",
		})
		return
	case user := <-resultCh:
		if user == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Can not retrieve user from database",
			})
			return
		}
		// Database operation completed before the context timeout
		c.JSON(http.StatusOK, &userInfoResponse{
			Username:    user.Username,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			PhoneNumber: user.PhoneNumber,
			Birthday:    user.Birthday,
		})
		return
	}
}
