package model

import (
	"gorm.io/gorm"
)

type AdditionalInfo struct {
	gorm.Model
	UserId  uint   `json:"user_id" form:"user_id" uri:"user_id" gorm:"not null" binding:"required" message:"user_id can not be empty" `
	OrderId uint   `json:"order_id" form:"order_id" uri:"order_id" gorm:"not null" binding:"required" message:"order_id can not be empty"`
	Message string `json:"message" form:"message" uri:"message" gorm:"size:128;not null" binding:"required" message:"message can not be empty" `
	Status  string `json:"status" form:"status" uri:"status" gorm:"size:128;not null" binding:"required" message:"status can not be empty" `
	Request string `json:"request" form:"request" uri:"request" gorm:"size:128;not null" message:"request can not be empty" `
}

type Message struct {
	gorm.Model
	UserId  uint   `json:"user_id" form:"user_id" uri:"user_id" gorm:"not null" binding:"required" message:"user_id can not be empty" `
	OrderId uint   `json:"order_id" form:"order_id" uri:"order_id" gorm:"not null" binding:"required" message:"order_id can not be empty"`
	Message string `json:"message" form:"message" uri:"message" gorm:"size:128;not null" binding:"required" message:"message can not be empty" `
	Status  string `json:"status" form:"status" uri:"status" gorm:"size:128;not null" binding:"required" message:"status can not be empty" `
}

type ProviderMsg struct {
	gorm.Model
	UserId  uint   `json:"user_id" form:"user_id" uri:"user_id" gorm:"not null" binding:"required" message:"user_id can not be empty" `
	Message string `json:"message" form:"message" uri:"message" gorm:"size:128;not null" binding:"required" message:"message can not be empty" `
	Status  string `json:"status" form:"status" uri:"status" gorm:"size:128;not null" `
}
