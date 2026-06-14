package config

type Captcha struct {
	ExpireSeconds int     `mapstructure:"expire_seconds"`
	Width         int     `mapstructure:"width"`
	Height        int     `mapstructure:"height"`
	Length        int     `mapstructure:"length"`
	MaxSkew       float64 `mapstructure:"max_skew"`
	DotCount      int     `mapstructure:"dot_count"`
}
