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

func (m *ServiceDao) ApproveServiceById(id uint) error {
	return m.Orm.Model(&model.Service{}).Where("id = ?", id).Update("status", "Approved").Error
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

func (m *ServiceDao) GetServiceByProviderId(iServiceListDTO *dto.ProviderServiceListDTO) ([]model.Service, int64, error) {
	var giServiceList []model.Service
	var nTotal int64

	err := m.Orm.Model(&model.Service{}).
		Scopes(Paginate(iServiceListDTO.Paginate)).
		Where("user_id = ?", iServiceListDTO.ProviderId).
		Find(&giServiceList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giServiceList, nTotal, err
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

func (m *ServiceDao) GetApprovedServiceList(iServiceListDTO *dto.ServiceListDTO) ([]model.Service, int64, error) {
	var giServiceList []model.Service
	var nTotal int64

	err := m.Orm.Model(&model.Service{}).
		Where("status = 'Approved'").
		Scopes(Paginate(iServiceListDTO.Paginate)).
		Find(&giServiceList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giServiceList, nTotal, err
}

func (m *ServiceDao) GetServiceListByCity(city string, iServiceListDTO *dto.ServiceListDTO) ([]model.Service, int64, error) {
	var giServiceList []model.Service
	var nTotal int64

	err := m.Orm.Model(&model.Service{}).
		Where("city = ? and status = 'Approved'", city).
		Scopes(Paginate(iServiceListDTO.Paginate)).
		Find(&giServiceList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giServiceList, nTotal, err
}

func (m *ServiceDao) GetApprovedServiceListByCategory(category string, iServiceListDTO *dto.ServiceListDTO) ([]model.Service, int64, error) {
	var giServiceList []model.Service
	var nTotal int64

	err := m.Orm.Model(&model.Service{}).
		Where("category = ? and status = 'Approved'", category).
		Scopes(Paginate(iServiceListDTO.Paginate)).
		Find(&giServiceList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giServiceList, nTotal, err
}

func (m *ServiceDao) GetServiceListByCityAndCategory(city string, category string, iServiceListDTO *dto.ServiceListDTO) ([]model.Service, int64, error) {
	var giServiceList []model.Service
	var nTotal int64

	err := m.Orm.Model(&model.Service{}).
		Where("city = ? AND category = ?", city, category).
		Scopes(Paginate(iServiceListDTO.Paginate)).
		Find(&giServiceList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giServiceList, nTotal, err
}
