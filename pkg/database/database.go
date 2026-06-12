// Package database mysql 数据库初始化
package database

import (
	"BlogServer/pkg/config"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg config.DB) *gorm.DB {
	var dialector gorm.Dialector
	switch cfg.Source {
	default:
		dialector = mysql.Open(cfg.DSN())
	}

	gormCfg := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成外键约束
	}
	if cfg.Debug {
		gormCfg.Logger = logger.Default.LogMode(logger.Info) // 打印 SQL
	}

	db, err := gorm.Open(dialector, gormCfg)
	if err != nil {
		zap.S().Fatalw("数据库连接失败", "error", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Fatalw("获取 sql.DB 失败", "error", err)
	}

	sqlDB.SetMaxIdleConns(10)                  // 空闲连接最大数量：峰值过后保留10个连接待命，防止空闲连接占用过多数据库资源
	sqlDB.SetMaxOpenConns(100)                 // 总连接数上限：最多同时存在100个连接（使用中+空闲），超出的请求阻塞等待，保护数据库不被压垮
	sqlDB.SetConnMaxLifetime(time.Hour)        // 连接最大存活时间：单个连接创建后最多使用1小时，到期强制关闭重建，防止网络设备超时断开或数据库会话僵死
	sqlDB.SetConnMaxIdleTime(30 * time.Minute) // 空闲连接超时：连接空闲超过30分钟强制关闭，避免维持长期不用的连接，减少数据库端会话维护开销

	//修改数据库默认字符集
	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
	//if !cfg.IsEmpty() {
	//	err = db.Use(dbresolver.Register(dbresolver.Config{
	//		// 指定主库：负责写操作
	//		Sources: []gorm.Dialector{mysql.Open(cfg.DSN())},
	//		// 指定从库列表：负责读操作
	//		Replicas: []gorm.Dialector{mysql.Open(cfg.DSN())},
	//		// 负载均衡策略：从库之间（RandomPolicy 随机 / RoundRobinPolicy 轮询）
	//		Policy: dbresolver.RandomPolicy{},
	//		// 调试时开启：打印 SQL 路由到主库还是从库
	//		TraceResolverMode: true,
	//	}))
	//
	//	if err != nil {
	//		zap.S().Fatalw("读写配置出错", "error", err)
	//	}
	//}

	zap.S().Infow("数据库连接成功",
		"source", cfg.Source,
		"host", cfg.Host,
		"port", cfg.Port,
		"db_name", cfg.DBName,
	)

	return db
}
