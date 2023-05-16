package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserId    uint   `json:"user_id" form:"user_id" gorm:"size:128;not null"`
	ServiceId uint   `json:"service_id" form:"service_id" gorm:"size:128;not null" uri:"service_id"`
	Content   string `json:"content" form:"content" gorm:"size:128;not null"`
	Rating    uint   `json:"rating" form:"rating" gorm:"size:128;not null" uri:"rating"`
}
