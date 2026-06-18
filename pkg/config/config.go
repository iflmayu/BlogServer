package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	System  System  `mapstructure:"system"`
	Log     Log     `mapstructure:"log"`
	DB      DB      `mapstructure:"db"`
	Jwt     Jwt     `mapstructure:"jwt"`
	Redis   Redis   `mapstructure:"redis"`
	Upload  Upload  `mapstructure:"upload"`
	Storage Storage `mapstructure:"storage"`
	Captcha Captcha `mapstructure:"captcha"`
	Email   Email   `mapstructure:"email"`
}

func LoadConfig(configPath string) (cfg *Config) {
	// 加载 .env 文件（本地开发用）
	_ = godotenv.Load()

	v := viper.New()

	pwd, _ := os.Getwd()
	fmt.Println("当前工作目录:", pwd)

	if configPath != "" {
		v.SetConfigFile(configPath) // 直接配置指定文件
	} else {
		v.SetConfigName("config") // 配置文件名（不带扩展名）
		v.SetConfigType("yaml")   // 明确指定类型
		v.AddConfigPath("./")     // 第一个搜索路径
		v.AddConfigPath("../../") // 第二个搜索路径（备用）
	}

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取yaml配置文件失败: %s", err))
	}

	cfg = new(Config)
	if err := v.Unmarshal(cfg); err != nil {
		panic(fmt.Errorf("解析yaml配置文件失败，配置文件格式错误: %s\n", err))
	}

	// 用环境变量覆盖敏感配置
	overrideFromEnv(cfg)

	fmt.Printf("配置文件 %s 加载成功！\n", configPath)

	return
}

func overrideFromEnv(cfg *Config) {
	if v := os.Getenv("BLOG_JWT_SECRET"); v != "" {
		cfg.Jwt.Secret = v
	}
	if v := os.Getenv("BLOG_EMAIL_PASSWORD"); v != "" {
		cfg.Email.Password = v
	}
	if v := os.Getenv("BLOG_QINIU_SECRET_KEY"); v != "" {
		cfg.Storage.Qiniu.SecretKey = v
	}
}
