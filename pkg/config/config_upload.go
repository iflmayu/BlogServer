package config

type Upload struct {
	MaxSize      int      `mapstructure:"max_size"`
	AllowedTypes []string `mapstructure:"allowed_types"`
	UploadDir    string   `mapstructure:"upload_dir"`
}
