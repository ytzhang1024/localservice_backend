package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerId   uint   `json:"customer_id" form:"customer_id" gorm:"size:128;not null" uri:"customer_id"`
	ServiceId    uint   `json:"service_id" form:"service_id" gorm:"size:128;not null" uri:"service_id"`
	Requirements string `json:"requirements" form:"requirements" gorm:"size:128"`
	Status       string
}
