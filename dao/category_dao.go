package dao

import (
	"6251/localservice/model"
)

type CategoryDao struct {
	BaseDao
}

var categoryDao *CategoryDao

func NewCategoryDao() *CategoryDao {
	if categoryDao == nil {
		categoryDao = &CategoryDao{NewBaseDao()}
	}

	return categoryDao
}

func (m *CategoryDao) GetCategoryById(id uint) (model.Category, error) {
	var iCategory model.Category
	err := m.Orm.First(&iCategory, id).Error
	return iCategory, err
}

// Query all categories
func (m *CategoryDao) GetCategoryList() ([]model.Category, error) {
	var iCategory []model.Category
	err := m.Orm.Find(&iCategory).Error
	return iCategory, err
}
