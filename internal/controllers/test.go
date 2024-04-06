package controllers

import (
	"github.com/gin-gonic/gin"
	"go-example/internal/services"
)

func Test(c *gin.Context) {
	code := c.Param("code")
	result, err := services.Test(code)
	RenderJSON(c, result, err)
}
