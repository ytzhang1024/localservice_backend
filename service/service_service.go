package service

import (
	"6251/localservice/dao"
	"6251/localservice/model"
	"6251/localservice/service/dto"
	"errors"
)

type ServiceService struct {
	BaseService
	Dao *dao.ServiceDao
}

var serviceService *ServiceService

func NewServiceService() *ServiceService {
	if serviceService == nil {
		serviceService = &ServiceService{
			Dao: dao.NewServiceDao(),
		}
	}
	return serviceService
}

func (m *ServiceService) AddService(iServiceAddDTO *dto.ServiceAddDTO) error {
	return m.Dao.AddService(iServiceAddDTO)
}

func (m *ServiceService) UpdateService(iServiceUpdateDTO *dto.ServiceUpdateDTO) error {
	if iServiceUpdateDTO.ID == 0 {
		return errors.New("Invalid Service ID")
	}

	return m.Dao.UpdateService(iServiceUpdateDTO)
}

func (m *ServiceService) GetServiceById(iCommonIDDTO *dto.CommonIDDTO) (model.Service, error) {
	return m.Dao.GetServiceById(iCommonIDDTO.ID)
}

func (m *ServiceService) GetServiceList(iServiceListDTO *dto.ServiceListDTO) ([]model.Service, int64, error) {
	return m.Dao.GetServiceList(iServiceListDTO)
}

func (m *ServiceService) DeleteServiceById(iCommonIDDTO *dto.CommonIDDTO) error {
	return m.Dao.DeleteServiceById(iCommonIDDTO.ID)
}
