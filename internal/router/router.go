package router

import (
	"github.com/gin-gonic/gin"
	"go-example/internal/controllers"
	"net/http"
)

func Run() {
	api := gin.New()
	api.GET("/test", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "code is cheap,show me the talking")

	})

	user := api.Group("user")
	user.GET("/:id", controllers.User)

	api.Run(":8089")
}
