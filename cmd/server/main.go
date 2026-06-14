package main

import (
	uploadDomain "BlogServer/internal/upload/domain"
	userDomain "BlogServer/internal/user/domain"
	"BlogServer/pkg/config"
	"BlogServer/pkg/database"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/logger"
	"BlogServer/pkg/redis"
	"BlogServer/pkg/router"
	"fmt"

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

	migrate(db)
	redis.Init(cfg.Redis)
	jwtService := jwt.NewService(cfg.Jwt.Secret, cfg.Jwt.Issuer, cfg.Jwt.Expire)
	tokenString, _ := jwtService.GenerateToken(jwt.Claims{
		UserID:   1,
		Username: "admin",
	})
	fmt.Println(tokenString)

	r := router.NewRouter(db, cfg, jwtService)
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
