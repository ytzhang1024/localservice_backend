package router

import (
	"6251/localservice/api"
	"github.com/gin-gonic/gin"
)

func InitAdditionalInfoRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		UpdateInfoApi := api.NewAdditionalInfoApi()
		rgAuthUpdateInfo := rgAuth.Group("update/info")
		{
			rgAuthUpdateInfo.POST("/add", UpdateInfoApi.AddAdditionalInfo)
			//rgAuthUpdateInfo.POST("/list", UpdateInfoApi.GetUpdateInfoList)
			//rgAuthUpdateInfo.GET("/:id", UpdateInfoApi.GetUpdateInfoById)
		}
	})
}
