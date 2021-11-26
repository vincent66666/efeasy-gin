package autoload

type Server struct {
	Http Http `mapstructure:"http" json:"http" yaml:"http"`
	Grpc Grpc `mapstructure:"grpc" json:"grpc" yaml:"grpc"`
}

type Http struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Timeout string `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}

type Grpc struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Timeout string `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}
