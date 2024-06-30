package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Variant struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	UUID        string `gorm:"not null" json:uuid`
	VariantName string `gorm:"not null" json:"variant_name" form:"variant_name" valid:"required~Variant name is required"`
	Quantity    int    `gorm:"not null"`
	ProductID   uint
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (b *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(b)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
