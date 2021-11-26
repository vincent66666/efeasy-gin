package global

import (
	"efeasy-gin/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Logger      *zap.Logger
	DB          *gorm.DB
	Trans       ut.Translator
	Redis       *redis.Client
	MinioClient *minio.Client
}

var App = new(Application)
