package main

import (
	"efeasy-gin/global"
	"efeasy-gin/initialize"
)

func main() {
	// 初始化配置
	global.App.ConfigViper = initialize.Viper()

	// 初始化日志
	global.App.Logger = initialize.Zap("Default")
	global.App.Logger.Info("Logger Init Success. ")

	// 初始化数据库
	global.App.DB = initialize.Gorm()
	global.App.Logger.Info("Gorm Init Success. ")

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			err := db.Close()
			if err != nil {
				panic(err)
			}
		}
	}()

	err := initialize.Validator("zh")
	if err != nil {
		panic(err)
	}
	global.App.Logger.Info("Validator Init Success. ")

	// 初始化Redis
	global.App.Redis = initialize.Redis()
	global.App.Logger.Info("Redis Init Success. ")

	// 7. 初始化minIO
	initialize.MinIO()

	// 启动服务器
	initialize.RunServer()
	global.App.Logger.Info("服务器启动成功. ")
}
