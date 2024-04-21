package user

import (
	"github.com/go-my-rss/go-my-rss/internal/service/user/dto"
	"gorm.io/gorm"
	"testing"
)

var userId uint

func TestCreateUser(t *testing.T) {
	create := dto.UserCreate{Username: "Jessy", Password: "123456", Email: "jessy@gmail.com", IsAdmin: true}
	id, err := CreateUser(&create)

	if err != nil {
		t.Error(err)
	}

	userId = id
}

func TestGetUser(t *testing.T) {
	get := dto.UserGet{Username: "Jessy", Password: "123456", Email: "jessy@gmail.com", IsAdmin: true}
	res, err := GetUser(&get)

	if err != nil {
		t.Error(err)
	}

	if res.Username != "Jessy" {
		t.Errorf("GetUser failed")
	}
}

func TestDeleteUser(t *testing.T) {
	userDelete := dto.UserDelete{Model: gorm.Model{ID: userId}}
	count, err := DeleteUser(&userDelete)
	if err != nil {
		t.Error(err)
	}

	if count == 0 {
		t.Errorf("DeleteUser failed")
	}
}
