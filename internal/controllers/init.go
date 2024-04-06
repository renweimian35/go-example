package controllers

import (
	"github.com/gin-gonic/gin"
	commonerrors "go-example/pkg/jerror"
	"net/http"
)

func RenderJSON(c *gin.Context, resp interface{}, err error) {
	errPayload := buildErrPayloadAndLog(err, true)
	if errPayload.Code == http.StatusOK {
		c.JSON(errPayload.Code, buildResponse(resp))
	} else {
		c.JSON(http.StatusInternalServerError, errPayload)
	}
	c.Next()
}

// nolint
func buildErrPayloadAndLog(err error, shouldLog bool) *commonerrors.ErrPayload {
	errPayload := commonerrors.NewErrPayload(err)
	if shouldLog {
		// 根据状态码打印日志
		if errPayload.Code < 400 {
			// 不需要打印日志
		} else if errPayload.Code >= 500 && errPayload.Code < 600 {
			//log2.ErrorDetails(err)
			// 对客户端隐藏详细信息
			errPayload.Message = ""
		} else {
			// 客户端错误 & 自定义错误，Warn
			//	log2.WarnDetails(err)
		}
	}
	return errPayload
}

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func buildResponse(data any) Resp {
	return Resp{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}
