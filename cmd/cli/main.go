package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"BlogServer/internal/user/domain"
	"BlogServer/internal/user/repo"
	"BlogServer/pkg/config"
	"BlogServer/pkg/database"
	"BlogServer/pkg/logger"
	"BlogServer/pkg/utils/pwd"

	"go.uber.org/zap"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create-admin":
		createAdmin()
	default:
		printUsage()
		os.Exit(1)
	}
}

func createAdmin() {
	cmd := flag.NewFlagSet("create-admin", flag.ExitOnError)
	username := cmd.String("username", "", "管理员用户名")
	email := cmd.String("email", "", "管理员邮箱")
	password := cmd.String("password", "", "管理员密码")

	if err := cmd.Parse(os.Args[2:]); err != nil {
		fmt.Println("参数解析失败:", err)
		os.Exit(1)
	}

	if *username == "" || *email == "" || *password == "" {
		fmt.Println("错误: username、email、password 都不能为空")
		cmd.Usage()
		os.Exit(1)
	}

	cfg := config.LoadConfig("config.yaml")
	logger.InitLog(cfg.Log)
	defer func() {
		_ = zap.L().Sync()
	}()

	db := database.InitDB(cfg.DB)
	userRepo := repo.NewUserRepo(db)
	ctx := context.Background()

	// 检查用户名/邮箱是否已存在
	exists, err := userRepo.ExistsUsername(ctx, *username)
	if err != nil {
		fmt.Println("检查用户名失败:", err)
		os.Exit(1)
	}
	if exists {
		fmt.Println("错误: 用户名已存在")
		os.Exit(1)
	}

	exists, err = userRepo.ExistsEmail(ctx, *email)
	if err != nil {
		fmt.Println("检查邮箱失败:", err)
		os.Exit(1)
	}
	if exists {
		fmt.Println("错误: 邮箱已存在")
		os.Exit(1)
	}

	hashedPassword, err := pwd.HashPassword(*password)
	if err != nil {
		fmt.Println("密码加密失败:", err)
		os.Exit(1)
	}

	admin := &domain.User{
		Username: *username,
		Email:    *email,
		Password: hashedPassword,
		Role:     domain.RoleAdmin,
	}

	if err := userRepo.Create(ctx, admin); err != nil {
		fmt.Println("创建管理员失败:", err)
		os.Exit(1)
	}

	fmt.Printf("管理员 %s 创建成功\n", *username)
}

func printUsage() {
	fmt.Println("BlogServer CLI 工具")
	fmt.Println()
	fmt.Println("用法:")
	fmt.Println("  go run cmd/cli/main.go create-admin --username=<用户名> --email=<邮箱> --password=<密码>")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  go run cmd/cli/main.go create-admin --username=admin --email=admin@example.com --password=admin")
}
