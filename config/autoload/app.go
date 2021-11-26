package autoload

type App struct {
	Env string `mapstructure:"env" json:"env" yaml:"env"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}
