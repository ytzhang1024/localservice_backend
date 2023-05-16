package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerId    uint   `json:"customer_id" form:"customer_id" gorm:"size:128;not null" uri:"customer_id"`
	ServiceId     uint   `json:"service_id" form:"service_id" gorm:"size:128;not null" uri:"service_id"`
	ProviderId    uint   `json:"provider_id" form:"provider_id" gorm:"size:128;not null" uri:"provider_id"`
	Description   string `json:"description" form:"description" gorm:"size:128"`
	Status        string `json:"status" form:"status" gorm:"size:128" uri:"status"`
	ServiceTitle  string `json:"service_title" form:"service_title" gorm:"size:128" uri:"service_title"`
	CustomerName  string `json:"customer_name" form:"customer_name" gorm:"size:128" uri:"customer_name"`
	CustomerEmail string `json:"customer_email" form:"customer_email" gorm:"size:128" uri:"customer_email"`
	CustomerPhone string `json:"customer_phone" form:"customer_phone" gorm:"size:128" uri:"customer_phone"`
	Postcode      string `json:"postcode" form:"postcode" gorm:"size:128" uri:"postcode"`
	Address       string `json:"address" form:"address" gorm:"size:128" uri:"address"`
	City          string `json:"city" form:"city" gorm:"size:128" uri:"city"`
	Date          string `json:"date" form:"date" gorm:"size:128" uri:"date"`
}
