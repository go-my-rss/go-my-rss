package structs

import "gorm.io/gorm"

type UserPermissionRel struct {
	UserId       int
	PermissionId int
	gorm.Model
}
