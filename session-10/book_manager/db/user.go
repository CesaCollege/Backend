package db

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username    string `gorm:"varchar(50), unique"`
	FirstName   string `gorm:"varchar(50)"`
	LastName    string `gorm:"varchar(50)"`
	PhoneNumber string `gorm:"varchar(30), unique"`
	Birthday    time.Time
	Password    string `gorm:"varchar(50)"`
}

func (gdb *GormDB) CreateNewUser(u *User) error {
	// Encrypting the user password
	if pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 4); err != nil {
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

func (gdb *GormDB) GetUserByUsername(username string) (*User, error) {
	var user User
	err := gdb.db.Where(&User{Username: username}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
