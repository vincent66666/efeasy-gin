package initialize

import (
	"efeasy-gin/utils"
	"go.uber.org/zap"
)

func Zap(group string) *zap.Logger {
	// 初始化 zap
	return utils.Logger(group)
}
