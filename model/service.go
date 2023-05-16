package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	UserId               int    `json:"user_id" form:"user_id" binding:"required" message:"user_id can not be empty"`
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
