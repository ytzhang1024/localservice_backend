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
			rgAuthUpdateInfo.POST("/message/add", UpdateInfoApi.AddMessage)
			rgAuthUpdateInfo.POST("/provider/message/add", UpdateInfoApi.AddProviderMsg)
			rgAuthUpdateInfo.GET("/find_message", UpdateInfoApi.GetMessageListByUserId)
			rgAuthUpdateInfo.GET("/find_provider_message", UpdateInfoApi.GetProviderMsgListByUserId)
			rgAuthUpdateInfo.GET("/find_info", UpdateInfoApi.GetAddiInfoListByUserId)
			rgAuthUpdateInfo.GET("/find_info_pending", UpdateInfoApi.GetAddiInfoListByUserIdAndStatus)
			rgAuthUpdateInfo.GET("/find_msg_pending", UpdateInfoApi.GetMsgByUserIdAndStatus)
			rgAuthUpdateInfo.PATCH("/update_info_status/:id", UpdateInfoApi.UpdateInfoStatus)
			rgAuthUpdateInfo.PATCH("/update_msg_status/:id", UpdateInfoApi.UpdateMsgStatus)
			//rgAuthUpdateInfo.POST("/list", UpdateInfoApi.GetUpdateInfoList)
			//rgAuthUpdateInfo.GET("/:id", UpdateInfoApi.GetUpdateInfoById)
		}
	})
}
