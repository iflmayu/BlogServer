package config

import "fmt"

type System struct {
	Ip      string `mapstructure:"ip"`
	Port    uint   `mapstructure:"port"`
	Env     string `mapstructure:"env"`
	GinMode string `mapstructure:"gin_mode"`
}

func (s *System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}
