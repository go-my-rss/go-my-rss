package dto

import "gorm.io/gorm"

type UserCreate struct {
	Username string
	Password string
	Email    string
	IsAdmin  bool
}

type UserGet struct {
	gorm.Model
	Username string
	Password string
	Email    string
	IsAdmin  bool
}

type UserUpdate struct {
	ID       uint
	UserName string
	Password string
	Email    string
	IsAdmin  bool
}

type UserDelete struct {
	gorm.Model
}

type UserPage struct {
}
