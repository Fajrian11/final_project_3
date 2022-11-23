package models

import (
	"errors"
	"strconv"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title" form:"title" valid:"required~Your Title is required"`
	Description string `gorm:"not null" json:"description" form:"description" valid:"required~Your Description is required"`
	Status      bool   `gorm:"not null" json:"status" form:"status" `
	UserID      uint
	User        *User
	CategoryID  uint `gorm:"not null" json:"category_id" form:"category_id" valid:"required~Your category_id is required"`
	Categories  *Categories
}

func (t *Task) BeforeCreate() (err error) {
	_, errCreate := govalidator.ValidateStruct(t)
	status := strconv.FormatBool(t.Status)
	if errCreate != nil {
		err = errCreate
		return
	}
	if status != "true" && status != "false" {
		err = errors.New("Status Hanya Boleh diisi dengan True or False")
		return err
	}
	err = nil
	return
}

// func (t *Task) BeforeUpdate() (err error) {
// 	_, errUpdate := govalidator.ValidateStruct(t)
// 	if errUpdate != nil {
// 		err = errUpdate
// 		return
// 	}
// 	if t.Status != true && t.Status != false {
// 		err = errors.New("Status Hanya Boleh diisi dengan True or False")
// 		return err
// 	}
// 	err = nil
// 	return
// }
