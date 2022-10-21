package models

import (
	"scalable-web-services-golang-hacktiv8/module_10/JWT/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	// Auto ada id, created, dan update atnya
	GormModel
	FullName string    `gorm:"not null" form:"full_name" json:"full_name" valid:"required~Full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Email is not valid"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password must be at least 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

// * Gorm hooks BeforeCreate
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	// Hash password
	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}
