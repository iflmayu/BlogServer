package config

import (
	"fmt"
	"os"

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

func LoadConfig(configPath string) (c *Config) {
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

	c = new(Config)
	if err := v.Unmarshal(c); err != nil {
		panic(fmt.Errorf("解析yaml配置文件失败，配置文件格式错误: %s\n", err))
	}

	fmt.Printf("配置文件 %s 加载成功！\n", configPath)

	return
}
