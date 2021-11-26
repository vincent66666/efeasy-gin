package response

import (
	"efeasy-gin/global/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`    // 自定义错误码
	Data    interface{} `json:"data"`    // 数据
	Message string      `json:"message"` // 信息
}

// Success 响应成功 ErrorCode 为 0 表示成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		response.Enum.Success.Code,
		data,
		response.Enum.Success.Message,
	})
}

// Fail 响应失败 ErrorCode 不为 200 表示失败
func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		nil,
		msg,
	})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, error response.Respond) {
	Fail(c, error.Code, error.Message)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, response.Enum.ValidateError.Code, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, response.Enum.BusinessError.Code, msg)
}

// TokenFail token验证失败
func TokenFail(c *gin.Context) {
	FailByError(c, response.Enum.TokenError)
}
