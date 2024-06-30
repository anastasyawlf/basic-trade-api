package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	UUID      string `gorm:"not null" json:uuid`
	Name      string `gorm:"not null" json:"product_name" form:"product_name" valid:"required~Product name is required"`
	ImageURL  string `gorm:"not null" json:"image_url" form:"image_url" valid:"required~Product image url is required"`
	AdminID   uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Variants  []Variant `gorm:"foreignkey:ProductID"`
}

func (b *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(b)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
