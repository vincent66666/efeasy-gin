package config

import (
	"efeasy-gin/config/autoload"
)

type Configuration struct {
	App      autoload.App      `mapstructure:"app" json:"app" yaml:"app"`
	Server   autoload.Server   `mapstructure:"server" json:"server" yaml:"server"`
	Logger   autoload.Logger   `mapstructure:"logger" json:"logger" yaml:"logger"`
	Database autoload.Database `mapstructure:"database" json:"database" yaml:"database"`
	Jwt      autoload.Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis    autoload.Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	MinIo    autoload.MinIo    `mapstructure:"minio" json:"minio" yaml:"minio"`
}
