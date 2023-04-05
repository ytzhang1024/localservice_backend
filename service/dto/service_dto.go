package dto

import (
	"6251/localservice/global/constants"
	"6251/localservice/model"
	"gorm.io/gorm"
)

type ServiceAddDTO struct {
	gorm.Model
	UserId               uint   `json:"userid" form:"userid" binding:"required" message:"userid can not be empty"`
	City                 string `json:"city" form:"city"`
	Title                string `json:"title" form:"title" binding:"required" message:"title can not be empty"`
	Description          string `json:"description" form:"description" binding:"required" message:"description can not be empty"`
	Prices               int    `json:"prices" form:"prices" binding:"required" message:"price can not be empty"`
	Address              string `json:"address" form:"address" binding:"required" message:"address can not be empty"`
	AreasCoverd          int    `json:"areas_coverd" form:"areas_coverd"`
	Category             string `json:"category" form:"category" binding:"required" message:"category can not be empty"`
	Availibility         string `json:"availibility" form:"availibility"`
	Mobile               string `json:"mobile" form:"mobile" binding:"required" message:"mobile can not be empty"`
	Photos               string `json:"photos" form:"photos"`
	LongitudeAndLatitude string `json:"longitude_latitude" form:"longitude_latitude"`
	Status               string
}

func (m *ServiceAddDTO) ConvertToModel(iService *model.Service) {
	iService.UserId = m.ID
	iService.City = m.City
	iService.Title = m.Title
	iService.Description = m.Description
	iService.Prices = m.Prices
	iService.Address = m.Address
	iService.AreasCoverd = m.AreasCoverd
	iService.Category = m.Category
	iService.Availibility = m.Availibility
	iService.Mobile = m.Mobile
	iService.Photos = m.Photos
	iService.LongitudeAndLatitude = m.LongitudeAndLatitude
	iService.Status = constants.PENDING
}

type ServiceUpdateDTO struct {
	gorm.Model
	City                 string `json:"city" form:"city"`
	Title                string `json:"title" form:"title" binding:"required" message:"title can not be empty"`
	Description          string `json:"description" form:"description" binding:"required" message:"description can not be empty"`
	Prices               int    `json:"prices" form:"prices" binding:"required" message:"price can not be empty"`
	Address              string `json:"address" form:"address" binding:"required" message:"address can not be empty"`
	AreasCoverd          int    `json:"areas_coverd" form:"areas_coverd"`
	Category             string `json:"category" form:"category" binding:"required" message:"category can not be empty"`
	Availibility         string `json:"availibility" form:"availibility"`
	Mobile               string `json:"mobile" form:"mobile" binding:"required" message:"mobile can not be empty"`
	Photos               string `json:"photos" form:"photos"`
	LongitudeAndLatitude string `json:"longitude_latitude" form:"longitude_latitude"`
}

func (m *ServiceUpdateDTO) ConvertToModel(iService *model.Service) {
	iService.City = m.City
	iService.Title = m.Title
	iService.Description = m.Description
	iService.Prices = m.Prices
	iService.Address = m.Address
	iService.AreasCoverd = m.AreasCoverd
	iService.Category = m.Category
	iService.Availibility = m.Availibility
	iService.Mobile = m.Mobile
	iService.Photos = m.Photos
	iService.LongitudeAndLatitude = m.LongitudeAndLatitude
	iService.Status = constants.PENDING
}

// ===============================================================================
// = Service List DTO
type ServiceListDTO struct {
	Paginate
}
