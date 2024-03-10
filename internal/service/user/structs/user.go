package structs

import "gorm.io/gorm"

type User struct {
	Username string
	Password string
	Email    string
	IsAdmin  bool
	gorm.Model
}
