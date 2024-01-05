package domain

import (
	"time"
)

type TMember struct {
	// gorm.Model
	UserId    string    `gorm:"column:USER_ID;primaryKey"`
	FirstName string    `gorm:"column:FIRST_NAME"`
	LastName  string    `gorm:"column:LAST_NAME"`
	Email     string    `gorm:"column:EMAIL"`
	CreateAt  time.Time `gorm:"column:CreateAt"`
}

func (TMember) TableName() string {
	return "t_member"
}
