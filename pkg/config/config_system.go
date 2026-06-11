package config

type System struct {
	Ip      string `mapstructure:"ip"`
	Port    uint   `mapstructure:"port"`
	Env     string `mapstructure:"env"`
	GinMode string `mapstructure:"gin_mode"`
}
