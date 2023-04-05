package dao

import (
	"6251/localservice/model"
	"6251/localservice/service/dto"
)

type ServiceDao struct {
	BaseDao
}

var serviceDao *ServiceDao

func NewServiceDao() *ServiceDao {
	if serviceDao == nil {
		serviceDao = &ServiceDao{NewBaseDao()}
	}

	return serviceDao
}

//func (m *UserDao) GetUserRoleById(email string) (int, error) {
//
//}

func (m *ServiceDao) AddService(iServiceAddDTO *dto.ServiceAddDTO) error {
	var iService model.Service
	iServiceAddDTO.ConvertToModel(&iService)

	err := m.Orm.Save(&iService).Error
	if err == nil {
		iServiceAddDTO.ID = iService.ID
	}

	return err
}

func (m *ServiceDao) DeleteServiceById(id uint) error {
	return m.Orm.Delete(&model.Service{}, id).Error
}

func (m *ServiceDao) UpdateService(iServiceUpdateDTO *dto.ServiceUpdateDTO) error {
	var iService model.Service

	m.Orm.First(&iService, iServiceUpdateDTO.ID)
	iServiceUpdateDTO.ConvertToModel(&iService)

	return m.Orm.Save(&iService).Error
}

func (m *ServiceDao) GetServiceById(id uint) (model.Service, error) {
	var iService model.Service
	err := m.Orm.First(&iService, id).Error
	return iService, err
}

func (m *ServiceDao) GetServiceList(iServiceListDTO *dto.ServiceListDTO) ([]model.Service, int64, error) {
	var giServiceList []model.Service
	var nTotal int64

	err := m.Orm.Model(&model.Service{}).
		Scopes(Paginate(iServiceListDTO.Paginate)).
		Find(&giServiceList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giServiceList, nTotal, err
}
