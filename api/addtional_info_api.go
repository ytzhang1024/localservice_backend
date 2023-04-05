package api

import (
	"6251/localservice/service"
	"6251/localservice/service/dto"
	"github.com/gin-gonic/gin"
	"net/http"
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

//查找所有关于此用户的Update Info
