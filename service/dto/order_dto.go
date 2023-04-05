package dto

import (
	"6251/localservice/model"
	"gorm.io/gorm"
)

type OrderDTO struct {
	gorm.Model
	CustomerId   uint   `json:"customer_id" form:"customer_id" gorm:"size:128;not null" uri:"customer_id"`
	ServiceId    uint   `json:"service_id" form:"service_id" gorm:"size:128;not null" uri:"service_id"`
	Requirements string `json:"requirements" form:"requirements" gorm:"size:128"`
	Status       string
}

func (m *OrderDTO) ConvertToModel(iOrder *model.Order) {
	iOrder.CustomerId = m.CustomerId
	iOrder.ServiceId = m.ServiceId
	iOrder.Status = m.Status
	iOrder.Requirements = m.Requirements
}

type OrderAddDTO struct {
	gorm.Model
	CustomerId   uint   `json:"customer_id" form:"customer_id" gorm:"size:128;not null" uri:"customer_id"`
	ServiceId    uint   `json:"service_id" form:"service_id" gorm:"size:128;not null" uri:"service_id"`
	Requirements string `json:"requirements" form:"requirements" gorm:"size:128"`
	Status       string
}

func (m *OrderAddDTO) ConvertToModel(iOrder *model.Order) {
	iOrder.CustomerId = m.CustomerId
	iOrder.ServiceId = m.ServiceId
	iOrder.Status = "Pending"
	iOrder.Requirements = m.Requirements
}

type OrderUpdateStateDTO struct {
	gorm.Model
	Status string `json:"status" form:"status" gorm:"size:128;not null";`
}

func (m *OrderUpdateStateDTO) ConvertToModel(iOrder *model.Order) {
	iOrder.Status = m.Status
}
