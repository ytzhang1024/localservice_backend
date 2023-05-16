package router

import (
	"6251/localservice/api"
	"github.com/gin-gonic/gin"
)

func InitServiceRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		serviceApi := api.NewServiceApi()
		rgPublicUser := rgPublic.Group("service").Use(func() gin.HandlerFunc {
			return func(ctx *gin.Context) {

			}
		}())
		{
			rgPublicUser.GET("/category", serviceApi.GetAllCategory)
			rgPublicUser.POST("/all_service", serviceApi.GetAllServiceList)
			rgPublicUser.POST("/list", serviceApi.GetAllApprovedServiceList)
			rgPublicUser.POST("/city", serviceApi.GetServiceListByCity)
			rgPublicUser.POST("/category", serviceApi.GetApprovedServiceListByCategory)
			rgPublicUser.POST("/city_n_cat", serviceApi.GetServiceListByCityAndCategory)
			rgPublicUser.GET("/:id", serviceApi.GetServiceById)
		}
		rgAuthService := rgAuth.Group("service")
		{
			rgAuthService.GET("/provider_service/:provider_id", serviceApi.GetServiceByProviderId)
			rgAuthService.PATCH("/approve/:id", serviceApi.ApproveServiceById)
			rgAuthService.POST("/add", serviceApi.AddService)
			rgAuthService.PUT("/:id", serviceApi.UpdateService)
			rgAuthService.DELETE("/delete/:id", serviceApi.DeleteServiceById)
		}
	})
}
