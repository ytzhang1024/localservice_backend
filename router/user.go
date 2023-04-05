package router

import (
	"6251/localservice/api"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user").Use(func() gin.HandlerFunc {
			return func(ctx *gin.Context) {

			}
		}())
		{
			rgPublicUser.POST("/login", userApi.Login)
			rgPublicUser.POST("/add", userApi.AddUser)
		}

		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.POST("/list", userApi.GetUserList)
			rgAuthUser.GET("/:id", userApi.GetUserById)
			rgAuthUser.PUT("/:id", userApi.UpdateUser)
			rgAuthUser.DELETE("/:id", userApi.DeleteUserById)
		}
	})
}
