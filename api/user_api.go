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

const (
	ERR_CODE_ADD_USER           = 10011
	ERR_CODE_GET_USER_BY_ID     = 10012
	ERR_CODE_GET_USER_LIST      = 10013
	ERR_CODE_UPDATE_USER        = 10014
	ERR_CODE_DELETE_USER        = 10015
	ERR_CODE_LOGIN              = 10016
	ERR_CODE_UPDATE_USER_AVATAR = 10017
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// @Tag User management
// @Summary User login
// @Description User login
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200 {string} string "Login Successful"
// @Failure 401 {string} string "Login Fail"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(c *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	iUser, token, err := m.Service.Login(iUserLoginDTO)
	if err == nil {
		err = service.SetLoginUserTokenToRedis(iUser.ID, token)
	}

	if err != nil {
		m.Fail(ResponseJson{
			Status: http.StatusUnauthorized,
			Code:   ERR_CODE_LOGIN,
			Msg:    err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Login Successful",
		Data: gin.H{
			"token":          token,
			"user_id":        iUser.ID,
			"user_email":     iUser.Email,
			"user_nick_name": iUser.NickName,
			"user_avatar":    iUser.Avatar,
			"user_mobile":    iUser.Mobile,
			"user_role":      iUser.Role,
		},
	})
}

func (m UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserAddDTO}).GetError(); err != nil {
		return
	}

	profile_pic := "upload/profile.jpg"
	iUserAddDTO.Avatar = profile_pic
	err := m.Service.AddUser(&iUserAddDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Register Successful",
		Data: iUserAddDTO,
	})
}

func (m UserApi) GetUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	iUser, err := m.Service.GetUserById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Data: iUser,
	})
}

func (m UserApi) GetUserList(c *gin.Context) {
	var iUserListDTO dto.UserListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserListDTO}).GetError(); err != nil {
		return
	}

	giUserList, nTotal, err := m.Service.GetUserList(&iUserListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data:  giUserList,
		Total: nTotal,
	})
}

func (m UserApi) GetProviderList(c *gin.Context) {
	var iProviderListDTO dto.UserListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iProviderListDTO}).GetError(); err != nil {
		return
	}

	giPrivoderList, nTotal, err := m.Service.GetProviderList(&iProviderListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data:  giPrivoderList,
		Total: nTotal,
	})
}

func (m UserApi) UpdateUser(c *gin.Context) {
	var iUserUpdateDTO dto.UserUpdateDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserUpdateDTO, BindAll: true}).GetError(); err != nil {
		return
	}

	// Check if avatar file exists
	if file, err := c.FormFile("avatar"); err == nil {
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
		iUserUpdateDTO.Avatar = dst
	}

	err := m.Service.UpdateUser(&iUserUpdateDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_UPDATE_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Update User Successful",
		Data: iUserUpdateDTO,
	})
}

//func (m UserApi) UpdateUser(c *gin.Context) {
//	var iUserUpdateDTO dto.UserUpdateDTO
//	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserUpdateDTO, BindAll: true}).GetError(); err != nil {
//		return
//	}
//
//	err := m.Service.UpdateUser(&iUserUpdateDTO)
//
//	if err != nil {
//		m.ServerFail(ResponseJson{
//			Code: ERR_CODE_UPDATE_USER,
//			Msg:  err.Error(),
//		})
//		return
//	}
//
//	m.OK(ResponseJson{
//		Code: http.StatusOK,
//		Msg:  "Update User Successful",
//		Data: iUserUpdateDTO,
//	})
//}

func (m UserApi) DeleteUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	err := m.Service.DeleteUserById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_DELETE_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Delete User Successful",
	})
}

func (m UserApi) ApproveServiceById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	err := m.Service.ApproveProviderById(&iCommonIDDTO)
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
