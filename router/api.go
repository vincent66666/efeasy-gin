package router

import (
	v1 "efeasy-gin/app/api/v1"
	"efeasy-gin/app/middleware"
	"efeasy-gin/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {

	// 设置跨域中间件
	//router.Use(middleware.CorsMiddleware.Cors())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/test", func(c *gin.Context) {
		time.Sleep(5*time.Second)
		c.String(http.StatusOK, "success")
	})



	router.POST("/v1/auth/login", v1.AuthApi.Login)
	authRouter := router.Group("/v1").Use(middleware.JwtMiddleware.JWTAuth(service.AppGuardName))
	{
		authRouter.POST("/users/create", v1.UserApi.UserCreate)
		authRouter.POST("/users", v1.UserApi.GetUserList)
		authRouter.POST("/auth/info", v1.AuthApi.Info)
		authRouter.POST("/auth/logout", v1.AuthApi.Logout)
	}

}
