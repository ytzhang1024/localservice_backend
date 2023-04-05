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
		}
		rgAuthService := rgAuth.Group("service")
		{
			rgAuthService.POST("/add", serviceApi.AddService)
		}
	})
}
