package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name     string `gorm:"column:name;type:varchar(100)" json:"name"`
	Password string `gorm:"column:password;type:varchar(32)" json:"password"`
	Phone    string `gorm:"column:phone;type:varchar(20)" json:"phone"`
	Email    string `gorm:"column:email;type:varchar(100)" json:"email"`
}

func (User) TableName() string {
	return "users"
}
