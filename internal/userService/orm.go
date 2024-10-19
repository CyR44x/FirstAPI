package userService

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Password string `json:"password"` // Наш сервер будет ожидать json c полем password
	Email    string `json:"email"`    //  Наш сервер будет ожидать json c полем email
}
