package config

type Log struct {
	App   string `mapstructure:"app"`
	Dir   string `mapstructure:"dir"`
	Level string `mapstructure:"level"`
}
