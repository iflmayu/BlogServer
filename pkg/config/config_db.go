package config

import "fmt"

type DB struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	Debug    bool   `mapstructure:"debug"` //打印全部日志
	Source   string `mapstructure:"source"`
}

func (d *DB) DSN() string {
	switch d.Source {
	default:
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			d.Username, d.Password, d.Host, d.Port, d.DBName)
	}
}

func (d *DB) IsEmpty() bool {
	return d.Host == "" || d.Port == 0 || d.DBName == ""
}
