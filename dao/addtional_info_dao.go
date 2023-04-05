package dao

import (
	"6251/localservice/model"
	"6251/localservice/service/dto"
)

type AdditionalInfoDao struct {
	BaseDao
}

var additionalInfoDao *AdditionalInfoDao

func NewAdditionalInfoDao() *AdditionalInfoDao {
	if additionalInfoDao == nil {
		additionalInfoDao = &AdditionalInfoDao{NewBaseDao()}
	}

	return additionalInfoDao
}

func (m *AdditionalInfoDao) AddAdditionalInfo(iAdditionalInfoDTO *dto.AdditionalInfoDTO) error {
	var iAdditionalInfo model.AdditionalInfo
	iAdditionalInfoDTO.ConvertToModel(&iAdditionalInfo)

	err := m.Orm.Save(&iAdditionalInfo).Error
	if err == nil {
		iAdditionalInfoDTO.ID = iAdditionalInfo.ID
	}

	return err
}
