package routes

import (
	"github.com/edu2801/first-golang-crud/api/controllers"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()
	v1 := router.Group("/v1")
	{
		v1.GET("/tweets", tweetController.FindAll)
		v1.POST("/tweet", tweetController.Create)
		v1.DELETE("/tweet/:id", tweetController.Delete)
	}

	return v1
}
