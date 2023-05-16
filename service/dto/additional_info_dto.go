package dto

import (
	"6251/localservice/model"
	"gorm.io/gorm"
)

type AdditionalInfoDTO struct {
	gorm.Model
	UserId  uint   `json:"user_id" form:"user_id" uri:"user_id" gorm:"not null" binding:"required" message:"user_id can not be empty" `
	OrderId uint   `json:"order_id" form:"order_id" uri:"order_id" gorm:"not null" binding:"required" message:"order_id can not be empty"`
	Message string `json:"message" form:"message" uri:"message" gorm:"size:128;not null" binding:"required" message:"message can not be empty" `
	Status  string `json:"status" form:"status" uri:"status" gorm:"size:128;not null" binding:"required" message:"status can not be empty" `
	Request string `json:"request" form:"request" uri:"request" gorm:"size:128;not null" message:"request can not be empty" `
}

func (m *AdditionalInfoDTO) ConvertToModel(iAdditionalInfo *model.AdditionalInfo) {
	iAdditionalInfo.UserId = m.UserId
	iAdditionalInfo.OrderId = m.OrderId
	iAdditionalInfo.Message = m.Message
	iAdditionalInfo.Status = m.Status
	iAdditionalInfo.Request = m.Request
}

type MessageDTO struct {
	gorm.Model
	UserId  uint   `json:"user_id" form:"user_id" uri:"user_id" gorm:"not null" binding:"required" message:"user_id can not be empty" `
	OrderId uint   `json:"order_id" form:"order_id" uri:"order_id" gorm:"not null" binding:"required" message:"order_id can not be empty"`
	Message string `json:"message" form:"message" uri:"message" gorm:"size:128;not null" binding:"required" message:"message can not be empty" `
	Status  string `json:"status" form:"status" uri:"status" gorm:"size:128;not null" binding:"required" message:"status can not be empty" `
}

func (m *MessageDTO) ConvertToModel(iMessage *model.Message) {
	iMessage.UserId = m.UserId
	iMessage.OrderId = m.OrderId
	iMessage.Message = m.Message
	iMessage.Status = m.Status
}

type ProviderMsgDTO struct {
	gorm.Model
	UserId  uint   `json:"user_id" form:"user_id" uri:"user_id" gorm:"not null" binding:"required" message:"user_id can not be empty" `
	Message string `json:"message" form:"message" uri:"message" gorm:"size:128;not null" binding:"required" message:"message can not be empty" `
	Status  string `json:"status" form:"status" uri:"status" gorm:"size:128;not null" `
}

func (m *ProviderMsgDTO) ConvertToModel(iProviderMsg *model.ProviderMsg) {
	iProviderMsg.UserId = m.UserId
	iProviderMsg.Message = m.Message
	iProviderMsg.Status = m.Status
}

type MessageListDTO struct {
	UserId uint `uri:"user_id" form:"user_id" message:"user_id can not be empty"`
	Paginate
}

type ProviderMsgListDTO struct {
	UserId uint `uri:"user_id" form:"user_id" message:"user_id can not be empty"`
	Paginate
}

type AdditionalInfoListDTO struct {
	UserId uint `uri:"user_id" form:"user_id" message:"user_id can not be empty"`
	Paginate
}
