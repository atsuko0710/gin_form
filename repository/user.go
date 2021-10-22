package repository

import (
	"gin_forum/models"

	"gorm.io/gorm"
)

func CreateUser(u models.User) error {
	return  models.CreateUser(u)
}

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) bool {
	_, err := models.FindByUsername(username);
	if err == gorm.ErrRecordNotFound {
		return true
	}
	if err != nil {
		return false
	}
	return false
}