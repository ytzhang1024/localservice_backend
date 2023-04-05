package service

import (
	"6251/localservice/dao"
	"6251/localservice/service/dto"
)

type AdditionalInfoService struct {
	BaseService
	Dao *dao.AdditionalInfoDao
}

var additionalInfoService *AdditionalInfoService

func NewAdditionalInfoService() *AdditionalInfoService {
	if additionalInfoService == nil {
		additionalInfoService = &AdditionalInfoService{
			Dao: dao.NewAdditionalInfoDao(),
		}
	}
	return additionalInfoService
}

func (m *AdditionalInfoService) AddAdditionalInfo(iAdditionalInfoDTO *dto.AdditionalInfoDTO) error {
	return m.Dao.AddAdditionalInfo(iAdditionalInfoDTO)
}
