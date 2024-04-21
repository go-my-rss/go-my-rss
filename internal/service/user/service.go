package user

import (
	"fmt"
	"github.com/go-my-rss/go-my-rss/internal/database"
	"github.com/go-my-rss/go-my-rss/internal/service/user/dto"
	"github.com/go-my-rss/go-my-rss/internal/service/user/structs"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = database.DB
}

func CreateUser(create *dto.UserCreate) (userId uint, err error) {
	user := &structs.User{}
	err = copier.Copy(user, create)
	if err != nil {
		return 0, fmt.Errorf("copy UserCreate to User failed: %w", err)
	}

	err = db.Create(&user).Error
	if err != nil {
		return 0, fmt.Errorf("create User failed: %w", err)
	}

	return user.ID, nil
}

func GetUser(get *dto.UserGet) (res *dto.UserGet, err error) {
	user := structs.User{}
	err = copier.Copy(&user, get)
	if err != nil {
		return nil, fmt.Errorf("copy UserGet to User failed: %w", err)
	}

	err = db.Where(user).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("get User failed: %w", err)
	}

	res = new(dto.UserGet)
	err = copier.Copy(res, &user)
	if err != nil {
		return nil, fmt.Errorf("copy User to UserGet failed: %w", err)
	}

	return
}

func DeleteUser(delete *dto.UserDelete) (deleteCount int, err error) {
	err = db.Unscoped().Delete(&structs.User{Model: gorm.Model{ID: delete.ID}}).Error
	if err != nil {
		return 0, fmt.Errorf("delete User failed: %w", err)
	}

	return 1, nil
}
