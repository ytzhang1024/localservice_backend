package service

import (
	"6251/localservice/dao"
	"6251/localservice/model"
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

func (m *AdditionalInfoService) AddMessage(iMessageDTO *dto.MessageDTO) error {
	return m.Dao.AddMessage(iMessageDTO)
}

func (m *AdditionalInfoService) AddProviderMsg(iMsgDTO *dto.ProviderMsgDTO) error {
	return m.Dao.AddProviderMessage(iMsgDTO)
}

func (m *AdditionalInfoService) GetMessageListByUserId(id int, iMessageListDTO *dto.MessageListDTO) ([]model.Message, int64, error) {
	return m.Dao.GetMessageListByUserId(id, iMessageListDTO)
}

func (m *AdditionalInfoService) GetProviderMsgListByUserId(id int, iMessageListDTO *dto.ProviderMsgListDTO) ([]model.ProviderMsg, int64, error) {
	return m.Dao.GetProviderMsgListByUserId(id, iMessageListDTO)
}

func (m *AdditionalInfoService) GetAdditionalInfoListByUserId(id int, iAdditionalInfoListDTO *dto.AdditionalInfoListDTO) ([]model.AdditionalInfo, int64, error) {
	return m.Dao.GetAdditionalInfoListByUserId(id, iAdditionalInfoListDTO)
}

func (m *AdditionalInfoService) UpdateInfoStatus(iCommonIDDTO *dto.CommonIDDTO, status string) error {
	return m.Dao.UpdateInfoStatus(iCommonIDDTO.ID, status)
}

func (m *AdditionalInfoService) UpdateMsgStatus(iCommonIDDTO *dto.CommonIDDTO, status string) error {
	return m.Dao.UpdateMsgStatus(iCommonIDDTO.ID, status)
}

func (m *AdditionalInfoService) GetAddiInfoListByUserIdAndStatus(id int, iAdditionalInfoListDTO *dto.AdditionalInfoListDTO) ([]model.AdditionalInfo, int64, error) {
	return m.Dao.GetAddiInfoListByUserIdAndStatus(id, iAdditionalInfoListDTO)
}

func (m *AdditionalInfoService) GetMsgByUserIdAndStatus(id int, iMsgListDTO *dto.MessageListDTO) ([]model.Message, int64, error) {
	return m.Dao.GetMsgByUserIdAndStatus(id, iMsgListDTO)
}
