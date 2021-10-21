package models

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

// func FindByParams(u *User) (*User, error) {
	
// }
