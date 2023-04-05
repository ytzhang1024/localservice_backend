package router

import (
	"6251/localservice/api"
	"github.com/gin-gonic/gin"
)

func InitReviewRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		reviewApi := api.NewReviewApi()
		rgPublicReview := rgPublic.Group("review").Use(func() gin.HandlerFunc {
			return func(ctx *gin.Context) {}
		}())
		{
			rgPublicReview.GET("/list/:service_id", reviewApi.GetReviewByServiceId)
		}
		rgAuthReview := rgAuth.Group("review")
		{
			rgAuthReview.POST("/add", reviewApi.AddReview)
			rgAuthReview.DELETE("/:id", reviewApi.DeleteReviewById)
		}
	})
}
