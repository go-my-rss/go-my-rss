package structs

import "gorm.io/gorm"

type Permission struct {
	Name string
	gorm.Model
}
