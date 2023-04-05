package dto

import (
	"6251/localservice/model"
)

type CategoryDTO struct {
	//gorm.Model
	ID   uint `gorm:"primarykey"`
	Name string
}

func (m *CategoryDTO) ConvertToModel(iCategory *model.Category) {
	iCategory.Name = m.Name
}
