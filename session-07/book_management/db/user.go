package db

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

// User contains data of each user account in the application
type User struct {
	gorm.Model
	Username    string `gorm:"varchar(50),unique"`
	FirstName   string `gorm:"varchar(50)"`
	LastName    string `gorm:"varchar(50)"`
	PhoneNumber string `gorm:"varchar(30),unique"`
	Password    string `gorm:"varchar(64)"`
}

// String returns a string representing the fields of User
func (u *User) String() string {
	return fmt.Sprintf("User{%+v}", *u)
}

// CreateNewUser creates a new user in the database using GormDB
func (gdb *GormDB) CreateNewUser(u *User) error {
	if pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 0); err != nil {
		return err
	} else {
		u.Password = string(pw)
	}

	// Check if there is a duplicate username
	var count int64
	gdb.db.Model(&User{}).Where(&User{Username: u.Username}).Count(&count)
	if count > 0 {
		return errors.New("this username is already taken")
	}

	return gdb.db.Create(u).Error
}

// GetUserByUsername returns the User entity that has the exact same username
func (gdb *GormDB) GetUserByUsername(username string) (*User, error) {
	var user User
	err := gdb.db.Where(&User{Username: strings.ToLower(username)}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
