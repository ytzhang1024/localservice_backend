package api

import (
	"6251/localservice/service"
	"6251/localservice/service/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdditionalInfoApi struct {
	BaseApi
	AdditionalInfo *service.AdditionalInfoService
}

func NewAdditionalInfoApi() AdditionalInfoApi {
	return AdditionalInfoApi{
		BaseApi:        NewBaseApi(),
		AdditionalInfo: service.NewAdditionalInfoService(),
	}
}

func (m AdditionalInfoApi) AddAdditionalInfo(c *gin.Context) {
	var iAdditionalInfoDTO dto.AdditionalInfoDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iAdditionalInfoDTO}).GetError(); err != nil {
		return
	}

	err := m.AdditionalInfo.AddAdditionalInfo(&iAdditionalInfoDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Additional Information Adding Successful",
		Data: iAdditionalInfoDTO,
	})
}

func (m AdditionalInfoApi) AddProviderMsg(c *gin.Context) {
	var iProviderMsgDTO dto.ProviderMsgDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iProviderMsgDTO}).GetError(); err != nil {
		return
	}

	err := m.AdditionalInfo.AddProviderMsg(&iProviderMsgDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Additional Provider Message Adding Successful",
		Data: iProviderMsgDTO,
	})
}

func (m AdditionalInfoApi) AddMessage(c *gin.Context) {
	var iMessageDTO dto.MessageDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iMessageDTO}).GetError(); err != nil {
		return
	}

	err := m.AdditionalInfo.AddMessage(&iMessageDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Additional Message Adding Successful",
		Data: iMessageDTO,
	})
}

func (m AdditionalInfoApi) GetMessageListByUserId(c *gin.Context) {
	var iMessageListDTO dto.MessageListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iMessageListDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	id_str := c.Query("user_id")
	id, _ := strconv.Atoi(id_str)

	giAllMessageList, nTotal, err := m.AdditionalInfo.GetMessageListByUserId(id, &iMessageListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code:  http.StatusOK,
		Data:  giAllMessageList,
		Total: nTotal,
	})
}

func (m AdditionalInfoApi) GetAddiInfoListByUserId(c *gin.Context) {
	var iAdditionalInfoListDTO dto.AdditionalInfoListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iAdditionalInfoListDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	id_str := c.Query("user_id")
	id, _ := strconv.Atoi(id_str)

	giAllAdditionalInfoList, nTotal, err := m.AdditionalInfo.GetAdditionalInfoListByUserId(id, &iAdditionalInfoListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code:  http.StatusOK,
		Data:  giAllAdditionalInfoList,
		Total: nTotal,
	})
}

func (m AdditionalInfoApi) UpdateInfoStatus(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	status := c.Query("status")
	println(status)

	err := m.AdditionalInfo.UpdateInfoStatus(&iCommonIDDTO, status)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Update Message Status Successful",
	})
}

func (m AdditionalInfoApi) UpdateMsgStatus(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	status := c.Query("status")
	println(status)

	err := m.AdditionalInfo.UpdateMsgStatus(&iCommonIDDTO, status)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Update Message Status Successful",
	})
}

func (m AdditionalInfoApi) GetAddiInfoListByUserIdAndStatus(c *gin.Context) {
	var iAdditionalInfoListDTO dto.AdditionalInfoListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iAdditionalInfoListDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	id_str := c.Query("user_id")
	id, _ := strconv.Atoi(id_str)

	giAllAdditionalInfoList, nTotal, err := m.AdditionalInfo.GetAddiInfoListByUserIdAndStatus(id, &iAdditionalInfoListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code:  http.StatusOK,
		Data:  giAllAdditionalInfoList,
		Total: nTotal,
	})
}

func (m AdditionalInfoApi) GetMsgByUserIdAndStatus(c *gin.Context) {
	var iMsgDTO dto.MessageListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iMsgDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	id_str := c.Query("user_id")
	id, _ := strconv.Atoi(id_str)

	giAllMsgList, nTotal, err := m.AdditionalInfo.GetMsgByUserIdAndStatus(id, &iMsgDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code:  http.StatusOK,
		Data:  giAllMsgList,
		Total: nTotal,
	})
}

func (m AdditionalInfoApi) GetProviderMsgListByUserId(c *gin.Context) {
	var iProviderMsgListDTO dto.ProviderMsgListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iProviderMsgListDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	id_str := c.Query("user_id")
	id, _ := strconv.Atoi(id_str)

	giAllProviderMsgList, nTotal, err := m.AdditionalInfo.GetProviderMsgListByUserId(id, &iProviderMsgListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code:  http.StatusOK,
		Data:  giAllProviderMsgList,
		Total: nTotal,
	})
}
