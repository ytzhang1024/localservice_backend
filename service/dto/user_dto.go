package dto

import "6251/localservice/model"

type UserLoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" message:"Email can not be empty"`
	Password string `json:"password" form:"password" binding:"required" message:"Password can not be empty"`
	Role     string `json:"role" form:"role"`
}

// ===============================================================================
// = Add user DTO
type UserAddDTO struct {
	ID       uint
	Email    string `json:"email" form:"email" binding:"required" message:"Email can not be empty"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"Password can not be empty"`
	NickName string `json:"nick_name" form:"nick_name"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Role     string `json:"role" form:"role" binding:"required" message:"Role can not be empty"`
}

func (m *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Email = m.Email
	iUser.NickName = m.NickName
	iUser.Avatar = m.Avatar
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
	iUser.Password = m.Password
	iUser.Role = m.Role
}

// ===============================================================================
// = Update user DTO
type UserUpdateDTO struct {
	ID       uint   `json:"id" form:"id" uri:"id"`
	Email    string `json:"email" form:"email"`
	NickName string `json:"nick_name" form:"nick_name"`
	Mobile   string `json:"mobile" form:"mobile"`
}

func (m *UserUpdateDTO) ConvertToModel(iUser *model.User) {
	iUser.ID = m.ID
	iUser.Email = m.Email
	iUser.NickName = m.NickName
	iUser.Mobile = m.Mobile
}

// ===============================================================================
// = User List DTO
type UserListDTO struct {
	Paginate
}
