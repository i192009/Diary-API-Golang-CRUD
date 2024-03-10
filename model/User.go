package model

import (
	Database "diaryApi/database"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255; not null; unigue" json:"username"`
	Password string `gorm:"size:255; not null" json:"-"`
	Entries  []Entry
}

func (user *User) Save() (*User, error) {
	err := Database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, err
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return err
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := Database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, err
}

func GetAllUsers() ([]User, error) {
	var user []User
	err := Database.Database.Find(&user).Error
	if err != nil {
		return user , err
	}
	return user , err

}
