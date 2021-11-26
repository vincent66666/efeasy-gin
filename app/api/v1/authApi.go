package v1

import (
	"efeasy-gin/app/request"
	"efeasy-gin/app/service"
	"efeasy-gin/utils"
	"efeasy-gin/utils/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authApi struct {
}

var AuthApi = new(authApi)

// Login 登录
func (authApi *authApi) Login(c *gin.Context) {
	var form request.LoginRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	err, user := service.UserService.Login(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	tokenData, err, _ := service.JwtService.CreateToken(service.AppGuardName, user)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, tokenData)
	return
}

// Info 用户信息
func (authApi *authApi) Info(c *gin.Context) {
	err, user := service.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
	return
}

// Logout 登出
func (authApi *authApi) Logout(c *gin.Context) {
	err := service.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	response.Success(c, nil)
	return
}

