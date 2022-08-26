package code

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code" binding:"required"`
	Data    interface{} `json:"data" binding:"required"`
	Msg     string      `json:"msg" binding:"required"`
	Success bool        `json:"success" binding:"required"`
}

func Result(code int, data interface{}, msg string, c *gin.Context, success bool) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
		success,
	})
}

// SuccessResponse 成功返回
func SuccessResponse(data interface{}, code int, c *gin.Context) {
	Result(code, data, Text(code), c, true)
}

// FailResponse 失败返回
func FailResponse(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, Text(code), c, false)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c, false)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c, false)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c, true)
}
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c, true)
}
