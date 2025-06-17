package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"primaryKey" json:"id"`
	//Email    string `json:"username" gorm:"unique;" validate:"required,email,min=6,max=32"`
	//Password string `json:"-" gorm:"type:text;" validate:"required,min=6"`
	Name  string `gorm:"type:varchar(20);not null" json:"name"`
	Phone string `gorm:"type:varchar(20);not null" json:"phone"`
}

//func (l User) Validate() error {
//	v := validator.New()
//	return v.Struct(l)
//}
