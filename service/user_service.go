package service

import (
	"6251/localservice/dao"
	"6251/localservice/global"
	"6251/localservice/global/constants"
	"6251/localservice/model"
	"6251/localservice/service/dto"
	"6251/localservice/utils"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}

	return userService
}

func SetLoginUserTokenToRedis(uid uint, token string) error {
	return global.RedisClient.Set(strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(uid)), -1), token, viper.GetDuration("jwt.tokenExpire")*time.Minute)
}

func (m *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, string, error) {
	var errResult error
	var token string

	iUser, err := m.Dao.GetUserByEmail(iUserDTO.Email)
	// Email or Password not correct
	if err != nil || !utils.CompareHashAndPassword(iUser.Password, iUserDTO.Password) {
		errResult = errors.New("Invalid Email Or Password")
	} else { // Login successful and generate token
		token, err = utils.GenerateToken(iUser.ID, iUser.Email)
		if err != nil {
			errResult = errors.New(fmt.Sprintf("Generate Token Error: %s", err.Error()))
		}
	}

	return iUser, token, errResult
}

func (m *UserService) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	if m.Dao.CheckEmailExist(iUserAddDTO.Email) {
		return errors.New("User Exist")
	}
	return m.Dao.AddUser(iUserAddDTO)
}

func (m *UserService) GetUserById(iCommonIDDTO *dto.CommonIDDTO) (model.User, error) {
	return m.Dao.GetUserById(iCommonIDDTO.ID)
}

func (m *UserService) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	return m.Dao.GetUserList(iUserListDTO)
}

func (m *UserService) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	if iUserUpdateDTO.ID == 0 {
		return errors.New("Invalid User ID")
	}

	return m.Dao.UpdateUser(iUserUpdateDTO)
}

func (m *UserService) DeleteUserById(iCommonIDDTO *dto.CommonIDDTO) error {
	return m.Dao.DeleteUserById(iCommonIDDTO.ID)
}
