package router

import (
	"6251/localservice/api"
	"github.com/gin-gonic/gin"
)

func InitOrderRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		orderApi := api.NewOrderApi()

		rgAuthOrder := rgAuth.Group("order").Use(func() gin.HandlerFunc {
			return func(ctx *gin.Context) {}
		}())
		{
			rgAuthOrder.POST("/add", orderApi.AddOrder)
			rgAuthOrder.GET("/my/:customer_id", orderApi.GetOrderByUserId)
			rgAuthOrder.GET("/service/:service_id", orderApi.GetOrderByServiceId)
			rgAuthOrder.PUT("/:id", orderApi.UpdateOrder)
		}
	})
}
