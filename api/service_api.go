package api

import (
	"6251/localservice/service"
	"6251/localservice/service/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServiceApi struct {
	BaseApi
	Service         *service.ServiceService
	CategoryService *service.CategoryService
}

func NewServiceApi() ServiceApi {
	return ServiceApi{
		BaseApi:         NewBaseApi(),
		Service:         service.NewServiceService(),
		CategoryService: service.NewCategoryService(),
	}
}

func (m ServiceApi) AddService(c *gin.Context) {
	var iServiceDTO dto.ServiceAddDTO

	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iServiceDTO}).GetError(); err != nil {
		return
	}

	//如果用户的角色是service provider
	//if () {}
	err := m.Service.AddService(&iServiceDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Service Adding Successful",
		Data: iServiceDTO,
	})
}

// Category List on Navigation bar
func (m ServiceApi) GetAllCategory(c *gin.Context) {
	if err := m.BuildRequest(BuildRequestOption{Ctx: c}).GetError(); err != nil {
		return
	}

	CategoryList, err := m.CategoryService.GetCategoryList()

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data: CategoryList,
	})
}
