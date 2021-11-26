package middleware

import (
	"efeasy-gin/global"
	"efeasy-gin/utils/response"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

type recoveryMiddleware struct {
}

var RecoveryMiddleware = new(recoveryMiddleware)

func (recoveryMiddleware *recoveryMiddleware) CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   global.App.Config.Logger.Default.RootDir + "/" + global.App.Config.Logger.Default.Filename,
			MaxSize:    global.App.Config.Logger.Default.MaxSize,
			MaxBackups: global.App.Config.Logger.Default.MaxBackups,
			MaxAge:     global.App.Config.Logger.Default.MaxAge,
			Compress:   global.App.Config.Logger.Default.Compress,
		},
		response.ServerError)
}
