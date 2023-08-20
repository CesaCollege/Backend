package authenticate

import (
	"bookman/db"
	"crypto/rand"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Auth struct {
	db                    *db.GormDB
	secretKey             []byte
	jwtExpirationDuration time.Duration
	logger                *logrus.Logger
}

func NewAuth(db *db.GormDB, jwtExpirationDuration time.Duration, logger *logrus.Logger) (*Auth, error) {
	// Create secret key
	secretKey, err := generateRandomKey()
	if err != nil {
		return nil, err
	}

	// Check database
	if db == nil {
		return nil, errors.New("database can not be nil")
	}

	return &Auth{
		db:                    db,
		secretKey:             secretKey,
		jwtExpirationDuration: jwtExpirationDuration,
		logger:                logger,
	}, nil
}

type Credentials struct {
	Username string
	Password string
}

type Token struct {
	TokenString string
}

type claims struct {
	jwt.MapClaims
	Username string `json:"username"`
}

func (a *Auth) Login(cred Credentials) (*Token, error) {
	// Check if the user exists
	account, err := a.db.GetUserByUsername(cred.Username)
	if err != nil {
		return nil, err
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(cred.Password))
	if err != nil {
		return nil, errors.New("the password is not correct")
	}

	// Create JWT token
	expirationTime := time.Now().Add(a.jwtExpirationDuration)
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		Username: cred.Username,
		MapClaims: jwt.MapClaims{
			"expired_at": expirationTime.Unix(),
		},
	})
	tokenString, err := tokenJWT.SignedString(a.secretKey)
	if err != nil {
		return nil, err
	}

	return &Token{TokenString: tokenString}, err
}

func (a *Auth) GetAccountByToken(token string) (*string, error) {
	// Handle empty token
	if token == "" {
		return nil, errors.New("access denied: the token is empty")
	}

	// Validate JWT token
	claim, err := a.checkToken(token)
	if err != nil {
		return nil, errors.New("access denied: the access token is not valid")
	}

	return &claim.Username, nil
}

func (a *Auth) checkToken(tokenStr string) (*claims, error) {
	c := &claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, c, func(token *jwt.Token) (interface{}, error) {
		return a.secretKey, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token")
		}

		a.logger.WithError(err).Warn("can not validate the token of the user")
		return nil, errors.New("bad error in validating user token")
	}

	if !tkn.Valid {
		return nil, errors.New("unauthorized")
	}

	return c, nil
}

func generateRandomKey() ([]byte, error) {
	jwtKey := make([]byte, 32)
	if _, err := rand.Read(jwtKey); err != nil {
		return nil, err
	}

	return jwtKey, nil
}
