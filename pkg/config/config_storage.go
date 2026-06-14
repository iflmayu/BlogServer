package config

type Storage struct {
	Type  string       `mapstructure:"type"`
	Local LocalStorage `mapstructure:"local"`
	Qiniu QiniuStorage `mapstructure:"qiniu"`
}

type LocalStorage struct {
	RootPath string `mapstructure:"root_path"`
	BaseURL  string `mapstructure:"base_url"`
}

type QiniuStorage struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
	Domain    string `mapstructure:"domain"`
}
