package models

import "gin_forum/config/mysql"

type User struct {
	BaseModel
	UserId   int64  `json:"user_id" gorm:"column:user_id;not null"`
	UserName string `json:"username" gorm:"column:username;not null" bindding:"required"`
	Password string `json:"password" gorm:"column:password;not null" bindding:"required"`
	Email    string `json:"email" gorm:"column:email;not null" bindding:"required"`
	Gender   int    `json:"gender" gorm:"column:gender;not null" bindding:"required"`
}

func (u *User) TableName() string {
	return TNUser()
}

// FindByUsername 根据用户名查找用户信息
func FindByUsername(username string) (*User, error) {
	u := &User{}
	d := mysql.Db.Where("username=?", username).First(&u)
	return u, d.Error
}

// CreateUser 创建用户
func CreateUser(u User) error {
	return mysql.Db.Create(&u).Error
}
