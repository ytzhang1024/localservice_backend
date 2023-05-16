package api

import (
	"6251/localservice/service"
	"6251/localservice/service/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
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

	iServiceDTO.Photos = "upload/logo.png"

	//如果用户的角色是service provider 才可以添加，前端校验
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

func (m ServiceApi) UpdateService(c *gin.Context) {
	var iServiceUpdateDTO dto.ServiceUpdateDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iServiceUpdateDTO, BindAll: true}).GetError(); err != nil {
		return
	}

	// Check if avatar file exists
	if file, err := c.FormFile("photo"); err == nil {
		// Check file type
		if !strings.HasSuffix(file.Filename, ".jpg") && !strings.HasSuffix(file.Filename, ".jpeg") && !strings.HasSuffix(file.Filename, ".png") {
			m.ServerFail(ResponseJson{
				Code: ERR_CODE_UPDATE_USER_AVATAR,
				Msg:  "invalid file type",
			})
			return
		}

		// Check file size
		if file.Size > 2<<20 { // 2MB
			m.ServerFail(ResponseJson{
				Code: ERR_CODE_UPDATE_USER_AVATAR,
				Msg:  "file size exceeds limit",
			})
			return
		}

		// Save file
		dst := fmt.Sprintf("upload/%d%s", time.Now().UnixNano(), file.Filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			m.ServerFail(ResponseJson{
				Code: ERR_CODE_UPDATE_USER_AVATAR,
				Msg:  err.Error(),
			})
			return
		}

		// Update avatar path in DTO
		iServiceUpdateDTO.Photos = dst
	}

	// Update the service
	err := m.Service.UpdateService(&iServiceUpdateDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: 500,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Service Update Successful",
		Data: iServiceUpdateDTO,
	})
}

//func (m ServiceApi) UpdateService(c *gin.Context) {
//	var iServiceUpdateDTO dto.ServiceUpdateDTO
//	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iServiceUpdateDTO, BindAll: true}).GetError(); err != nil {
//		return
//	}
//
//	// Parse the multipart form
//	err := c.Request.ParseMultipartForm(32 << 20) // limit to 32 MB
//	if err != nil {
//		m.ServerFail(ResponseJson{
//			Code: 500,
//			Msg:  "Failed to parse form: " + err.Error(),
//		})
//		return
//	}
//
//	// Get the uploaded files
//	files := c.Request.MultipartForm.File["pictures"]
//
//	// If there are files, process them
//	if len(files) > 0 {
//		// Create a slice to hold the file paths
//		picturePaths := make([]string, 0)
//
//		// Loop through each file
//		for _, file := range files {
//			// Generate a unique file name
//			fileName := uuid.New().String() + path.Ext(file.Filename)
//
//			// Save the file to disk
//			err = c.SaveUploadedFile(file, "upload/"+fileName)
//			if err != nil {
//				m.ServerFail(ResponseJson{
//					Code: 500,
//					Msg:  "Failed to save picture: " + err.Error(),
//					Data: iServiceUpdateDTO,
//				})
//				return
//			}
//
//			// Add the file path to the slice
//			picturePaths = append(picturePaths, fileName)
//		}
//
//		str := strings.Join(picturePaths, ", ")
//		// Set the picture paths in the DTO
//		iServiceUpdateDTO.Photos = str
//	}
//
//	// Update the service
//	err = m.Service.UpdateService(&iServiceUpdateDTO)
//
//	m.OK(ResponseJson{
//		Code: http.StatusOK,
//		Msg:  "Service Update Successful",
//		Data: iServiceUpdateDTO,
//	})
//}

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
		Code: 200,
		Data: CategoryList,
	})
}

// Get All Service Provider
func (m ServiceApi) GetAllServiceList(c *gin.Context) {
	var iAllServiceListDTO dto.ServiceListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iAllServiceListDTO}).GetError(); err != nil {
		return
	}

	giAllServiceList, nTotal, err := m.Service.GetServiceList(&iAllServiceListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code:  200,
		Data:  giAllServiceList,
		Total: nTotal,
	})
}

func (m ServiceApi) GetAllApprovedServiceList(c *gin.Context) {
	var iAllServiceListDTO dto.ServiceListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iAllServiceListDTO}).GetError(); err != nil {
		return
	}
	giAllServiceList, nTotal, err := m.Service.GetApprovedServiceList(&iAllServiceListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code:  http.StatusOK,
		Data:  giAllServiceList,
		Total: nTotal,
	})
}

func (m ServiceApi) GetServiceByProviderId(c *gin.Context) {
	var iProviderServiceListDTO dto.ProviderServiceListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iProviderServiceListDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	giProviderServiceListDTO, nTotal, err := m.Service.GetServiceByProviderId(&iProviderServiceListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code:  http.StatusOK,
		Data:  giProviderServiceListDTO,
		Total: nTotal,
	})
}

func (m ServiceApi) GetServiceById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	iService, err := m.Service.GetServiceById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Data: iService,
	})
}

func (m ServiceApi) DeleteServiceById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	err := m.Service.DeleteServiceById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_DELETE_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Delete Service Successful",
	})
}

func (m ServiceApi) ApproveServiceById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	err := m.Service.ApproveServiceById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Approve Service Successful",
	})
}

func (m ServiceApi) GetServiceListByCity(c *gin.Context) {
	var iAllServiceListDTO dto.ServiceListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iAllServiceListDTO}).GetError(); err != nil {
		return
	}

	city := c.Query("city")

	giAllServiceList, nTotal, err := m.Service.GetServiceListByCity(city, &iAllServiceListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code:  http.StatusOK,
		Data:  giAllServiceList,
		Total: nTotal,
	})
}

func (m ServiceApi) GetApprovedServiceListByCategory(c *gin.Context) {
	var iAllServiceListDTO dto.ServiceListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iAllServiceListDTO}).GetError(); err != nil {
		return
	}

	category := c.Query("category")

	giAllServiceList, nTotal, err := m.Service.GetApprovedServiceListByCategory(category, &iAllServiceListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code:  http.StatusOK,
		Data:  giAllServiceList,
		Total: nTotal,
	})
}

func (m ServiceApi) GetServiceListByCityAndCategory(c *gin.Context) {
	var iAllServiceListDTO dto.ServiceListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iAllServiceListDTO}).GetError(); err != nil {
		return
	}

	city := c.Query("city")
	category := c.Query("category")

	giAllServiceList, nTotal, err := m.Service.GetServiceListByCityAndCategory(city, category, &iAllServiceListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code:  http.StatusOK,
		Data:  giAllServiceList,
		Total: nTotal,
	})
}
