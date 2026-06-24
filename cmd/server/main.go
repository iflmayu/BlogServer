package main

import (
	uploadDomain "BlogServer/internal/upload/domain"
	userDomain "BlogServer/internal/user/domain"
	"BlogServer/pkg/captcha"
	"BlogServer/pkg/config"
	"BlogServer/pkg/database"
	"BlogServer/pkg/email"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/logger"
	"BlogServer/pkg/redis"
	"BlogServer/pkg/router"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig("config.yaml")
	logger.InitLog(cfg.Log)

	defer func() {
		_ = zap.L().Sync()
	}()
	//zap.S().Infow("测试")
	db := database.InitDB(cfg.DB)

	//migrate(db)
	redis.Init(cfg.Redis)
	captcha.Init(cfg.Captcha, redis.Client)

	jwtService := jwt.NewService(cfg.Jwt.Secret, cfg.Jwt.Issuer, cfg.Jwt.Expire)
	emailService := email.NewService(cfg.Email)

	//生成token
	//generateToken(jwtService)

	// 启动 web 服务
	r := router.NewRouter(db, cfg, jwtService, emailService)
	if err := r.Run(cfg.System.Addr()); err != nil {
		zap.S().Fatalw("服务器启动失败", "err", err)
	}
}

// 开发阶段自动迁移
func migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&userDomain.User{},
		&uploadDomain.Upload{},
	)
	if err != nil {
		zap.S().Fatalw("数据库迁移失败", "err", err)
		return
	}
	zap.S().Infow("数据库迁移成功")
}
