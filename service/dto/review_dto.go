package dto

import (
	"6251/localservice/model"
	"gorm.io/gorm"
)

type ReviewDTO struct {
	gorm.Model
	UserId    uint   `json:"user_id" form:"user_id" gorm:"size:128;not null"`
	ServiceId uint   `json:"service_id" form:"service_id" gorm:"size:128;not null" uri:"service_id"`
	Content   string `json:"content" form:"content" gorm:"size:128;not null"`
}

func (m *ReviewDTO) ConvertToModel(iReview *model.Review) {
	iReview.UserId = m.UserId
	iReview.ServiceId = m.ServiceId
	iReview.Content = m.Content
}
