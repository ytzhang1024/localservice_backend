package service

import (
	"6251/localservice/dao"
	"6251/localservice/model"
	"6251/localservice/service/dto"
)

type CategoryService struct {
	BaseService
	Dao *dao.CategoryDao
}

var categoryService *CategoryService

func NewCategoryService() *CategoryService {
	if categoryService == nil {
		categoryService = &CategoryService{
			Dao: dao.NewCategoryDao(),
		}
	}

	return categoryService
}

func (m *CategoryService) GetCategoryById(iCommonIDDTO *dto.CommonIDDTO) (model.Category, error) {
	return m.Dao.GetCategoryById(iCommonIDDTO.ID)
}

func (m *CategoryService) GetCategoryList() ([]model.Category, error) {
	return m.Dao.GetCategoryList()
}
