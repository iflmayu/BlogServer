package config

type Captcha struct {
	ExpireSeconds int `mapstructure:"expire_seconds"`
	Width         int `mapstructure:"width"`
	Height        int `mapstructure:"height"`
	Length        int `mapstructure:"length"`
}
