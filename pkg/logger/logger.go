package logger

import (
	"BlogServer/pkg/config"
	"fmt"

	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog(env config.Log) *zap.SugaredLogger {
	//确保日志目录存在
	if err := os.MkdirAll(env.Dir, os.ModePerm); err != nil {
		// 如果目录创建失败，直接停止程序，因为没有日志后续很难排查问题
		panic(fmt.Errorf("无法创建日志目录: " + err.Error()))
	}

	// Encoder 配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder, // 控制台彩色显示
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	// 创建不同的 Encoder
	// 控制台用 ConsoleEncoder（带颜色，易读）
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 文件通常建议用 JSONEncoder（结构化，方便分析）
	// 且输出到文件时，通常关闭颜色，避免乱码
	fileEncoderConfig := encoderConfig
	fileEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 文件去掉颜色
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)

	// 设置输出方向
	// 控制台输出
	consoleSyncer := zapcore.AddSync(os.Stdout)

	// 文件输出
	fileName := fmt.Sprintf("%s.log", env.App)
	lumberjackSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(env.Dir, fileName),
		MaxSize:    100,  // MB
		MaxBackups: 30,   // 备份数量
		MaxAge:     7,    // 保留天数
		Compress:   true, // 压缩
		LocalTime:  true, // 使用本地时间命名备份文件
	})

	// 使用 NewTee 组合多个 Core
	// Tee 可以为不同的输出指定不同的级别
	level := parseLevel(env.Level)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleSyncer, level), // 控制台全打
		zapcore.NewCore(fileEncoder, lumberjackSyncer, level), //
	)

	// 生成 Logger 并开启行号显示
	logger := zap.New(core, zap.AddCaller())

	// 替换全局 Logger
	zap.ReplaceGlobals(logger)

	return logger.Sugar()
}

func parseLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
