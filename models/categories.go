package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Type string `gorm:"not null" json:"type" form:"type" valid:"required~Your Type is required"`
	Task []Task `gorm:"foreignKey:CategoryID"`
}

func (c *Categories) BeforeCreate() (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (c *Categories) BeforeUpdate() (err error) {
	_, errUpdate := govalidator.ValidateStruct(c)
	if errUpdate != nil {
		err = errUpdate
		return
	}
	err = nil
	return
}
