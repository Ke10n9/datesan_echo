package model

import (
	"time"
)

type User struct {
	ID	uint      `json:"id"`
	Email	string	`json:"email" gorm:"type:varchar(255);unique;not null"`
	Password	string	`json:"password" gorm:"not null;default:0"`
	Created_at	time.Time	`json:"created_at"`
	Updated_at	time.Time	`json:"updated_at"`
}

func CreateUser(user *User) {
	DB.Select("Email", "Password").Create(user)
}

func FindUserById(id int) User {
	var user User
	DB.Where("ID", id).First(&user)
	return user
}

func FindUserByEmail(email string) User {
	var user User
	DB.Where("Email", email).First(&user)
	return user
}