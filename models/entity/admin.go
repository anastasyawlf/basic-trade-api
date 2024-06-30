package entity

import (
	"basic-trade/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UUID      string     `gorm:"not null" json:uuid`
	Name      string     `gorm:"not null" json:"name"" form:"name" valid:"required~Your full name is required"`
	Email     string     `gorm:"not null" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password  string     `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 4 characters"`
	CreatedAt *time.Time `json:"created_at, onitempty"`
	UpdatedAt *time.Time `json:"updated_at, onitempty"`
	Products  []Product  `constraint:OnUpdate:CASCADE,OnDelete:SET NULL;gorm:"foreignkey:AdminID" json:"products"`
}

func (u *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
