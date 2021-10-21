package models

import "time"

type BaseModel struct {
	Id         uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at;default:null" json:"create_time"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:null" json:"update_time"`
}
