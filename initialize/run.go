package initialize

import (
	"context"
	"efeasy-gin/global"
	"efeasy-gin/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setupRouter() *gin.Engine {
	GinRouter := gin.Default()

	// 前端项目静态资源
	GinRouter.StaticFile("/favicon.ico", "./storage/public/favicon.ico")
	// 其他静态资源
	GinRouter.Static("/public", "./storage/public")

	// 注册 api 分组路由
	apiGroup := GinRouter.Group("/api")
	router.SetApiGroupRoutes(apiGroup)

	return GinRouter
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	// 启动服务器
	srv := &http.Server{
		Addr:    global.App.Config.Server.Http.Addr + ":" + global.App.Config.Server.Http.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")


}
