package autoload

type Redis struct {
	Default RedisDefault `mapstructure:"default" json:"default" yaml:"default"`
}

type RedisDefault struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
