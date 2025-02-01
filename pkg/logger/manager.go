package logger

import (
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/logger/utils"
	"os"
	"sync"

	"github.com/ArtisanCloud/MediaXCore/pkg/logger/config"
	"github.com/ArtisanCloud/MediaXCore/pkg/logger/writer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	driver *zap.Logger
	config config.LogConfig
}

// 全局变量用来存储 Logger 实例
var (
	once     sync.Once
	instance *Logger
)

func NewLogger(config *config.LogConfig) *Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	var cores []zapcore.Core

	// Console 日志
	if config.Console {
		consoleCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), config.ParseLogLevel())
		cores = append(cores, consoleCore)
	}

	// 文件日志
	if config.File.Enable {
		if config.File.InfoFilePath == "" {
			config.File.InfoFilePath = "logs/info.log"
		}
		utils.EnsureFileExists(config.File.InfoFilePath)
		if config.File.ErrorFilePath == "" {
			config.File.ErrorFilePath = "logs/error.log"
		}
		utils.EnsureFileExists(config.File.ErrorFilePath)

		// 只记录 Info 及以上（但不包括 Error）级别的日志
		infoLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.InfoLevel && lvl < zapcore.ErrorLevel
		})
		infoFileCore := zapcore.NewCore(encoder, writer.NewFileWriter(config.File.InfoFilePath), infoLevelEnabler)
		cores = append(cores, infoFileCore)

		// 只记录 Error 及以上级别的日志
		errorLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
		errorFileCore := zapcore.NewCore(encoder, writer.NewFileWriter(config.File.ErrorFilePath), errorLevelEnabler)
		cores = append(cores, errorFileCore)
	}

	// Loki 日志
	if config.Loki.Enable {
		lokiCore := zapcore.NewCore(encoder, writer.NewLokiWriter(&config.Loki), config.ParseLogLevel())
		cores = append(cores, lokiCore)
	}

	// 组合多个日志 Core
	core := zapcore.NewTee(cores...)

	// 创建 Logger
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{
		driver: logger,
		config: *config,
	}
}

func GetLogger(c *config.LogConfig) *Logger {
	if c == nil {
		c = &config.LogConfig{
			Level:         "debug",
			Console:       true,
			UseJsonFormat: false,
			File: config.FileConfig{
				Enable: false,
			},
			Loki: config.LokiConfig{
				Enable: false,
			},
			HttpDebug: true,
			Debug:     true,
		}
	}

	once.Do(func() {
		// 仅在第一次调用时初始化
		instance = NewLogger(c)
	})
	return instance
}

// Info 输出 Info 日志
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.driver.Info(msg, fields...)
}

// Error 输出 Error 日志
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.driver.Error(msg, fields...)
}

// Debug 输出 Debug 日志
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.driver.Debug(msg, fields...)
}

// Warn 输出 Warn 日志
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.driver.Warn(msg, fields...)
}

// InfoF 格式化后输出 Info 日志
func (l *Logger) InfoF(format string, args ...interface{}) {
	l.driver.Info(fmt.Sprintf(format, args...))
}

// ErrorF 格式化后输出 Error 日志
func (l *Logger) ErrorF(format string, args ...interface{}) {
	l.driver.Error(fmt.Sprintf(format, args...))
}

// DebugF 格式化后输出 Debug 日志
func (l *Logger) DebugF(format string, args ...interface{}) {
	l.driver.Debug(fmt.Sprintf(format, args...))
}

// WarnF 格式化后输出 Warn 日志
func (l *Logger) WarnF(format string, args ...interface{}) {
	l.driver.Warn(fmt.Sprintf(format, args...))
}
