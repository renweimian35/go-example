package controllers

import (
	"github.com/gin-gonic/gin"
	"go-example/internal/services/user"
	"strconv"
)

func User(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		RenderJSON(c, nil, err)
		return
	}
	result, err := user.ById(i)
	RenderJSON(c, result, err)
}
