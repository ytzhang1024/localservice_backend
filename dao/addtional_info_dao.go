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

func (m *AdditionalInfoDao) AddMessage(iMessageDTO *dto.MessageDTO) error {
	var iMessage model.Message
	iMessageDTO.ConvertToModel(&iMessage)

	err := m.Orm.Save(&iMessage).Error
	if err == nil {
		iMessageDTO.ID = iMessage.ID
	}

	return err
}

func (m *AdditionalInfoDao) AddProviderMessage(iMsgDTO *dto.ProviderMsgDTO) error {
	var iMsg model.ProviderMsg
	iMsgDTO.ConvertToModel(&iMsg)

	err := m.Orm.Save(&iMsg).Error
	if err == nil {
		iMsgDTO.ID = iMsg.ID
	}

	return err
}

func (m *AdditionalInfoDao) GetMessageListByUserId(id int, iMessageListDTO *dto.MessageListDTO) ([]model.Message, int64, error) {
	var giMessageList []model.Message
	var nTotal int64

	err := m.Orm.Model(&model.Message{}).
		Scopes(Paginate(iMessageListDTO.Paginate)).
		Where("user_id = ?", id).
		Find(&giMessageList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giMessageList, nTotal, err
}

func (m *AdditionalInfoDao) GetProviderMsgListByUserId(id int, iMessageListDTO *dto.ProviderMsgListDTO) ([]model.ProviderMsg, int64, error) {
	var giMessageList []model.ProviderMsg
	var nTotal int64

	err := m.Orm.Model(&model.ProviderMsg{}).
		Scopes(Paginate(iMessageListDTO.Paginate)).
		Where("user_id = ?", id).
		Find(&giMessageList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giMessageList, nTotal, err
}

func (m *AdditionalInfoDao) GetAdditionalInfoListByUserId(id int, iAdditionalInfoListDTO *dto.AdditionalInfoListDTO) ([]model.AdditionalInfo, int64, error) {
	var giAdditionalInfoList []model.AdditionalInfo
	var nTotal int64

	err := m.Orm.Model(&model.AdditionalInfo{}).
		Scopes(Paginate(iAdditionalInfoListDTO.Paginate)).
		Where("user_id = ?", id).
		Find(&giAdditionalInfoList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giAdditionalInfoList, nTotal, err
}

func (m *AdditionalInfoDao) UpdateInfoStatus(id uint, status string) error {
	return m.Orm.Model(&model.AdditionalInfo{}).Where("id = ?", id).Update("status", status).Error
}

func (m *AdditionalInfoDao) UpdateMsgStatus(id uint, status string) error {
	return m.Orm.Model(&model.Message{}).Where("id = ?", id).Update("status", status).Error
}

func (m *AdditionalInfoDao) GetAddiInfoListByUserIdAndStatus(id int, iAdditionalInfoListDTO *dto.AdditionalInfoListDTO) ([]model.AdditionalInfo, int64, error) {
	var giAdditionalInfoList []model.AdditionalInfo
	var nTotal int64

	err := m.Orm.Model(&model.AdditionalInfo{}).
		Scopes(Paginate(iAdditionalInfoListDTO.Paginate)).
		Where("user_id = ? and status = 'Pending'", id).
		Find(&giAdditionalInfoList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giAdditionalInfoList, nTotal, err
}

func (m *AdditionalInfoDao) GetMsgByUserIdAndStatus(id int, iMsgListDTO *dto.MessageListDTO) ([]model.Message, int64, error) {
	var giMsgList []model.Message
	var nTotal int64

	err := m.Orm.Model(&model.Message{}).
		Scopes(Paginate(iMsgListDTO.Paginate)).
		Where("user_id = ? and status = 'Pending'", id).
		Find(&giMsgList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giMsgList, nTotal, err
}
