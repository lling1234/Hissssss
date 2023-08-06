package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type Config struct {
	Path     string `yaml:"path"`
	Size     int    `yaml:"size"`
	Keep     int    `yaml:"keep"`
	Backups  int    `yaml:"backups"`
	Compress bool   `yaml:"compress"`
}

func New(loggerConfig Config) *zap.Logger {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	fileEncode := zapcore.NewJSONEncoder(NewProductionEncoderConfig())
	fileSyncer := zapcore.AddSync(FileLogHook(loggerConfig))
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleDebugging := zapcore.Lock(os.Stdout)

	var cores []zapcore.Core

	cores = append(cores, zapcore.NewCore(fileEncode, fileSyncer, highPriority))
	cores = append(cores, zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
	core := zapcore.NewTee(cores...)
	logger := zap.New(core).WithOptions(zap.WithCaller(true))
	return logger
}

func FileLogHook(loggerConfig Config) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   loggerConfig.Path,
		MaxSize:    loggerConfig.Size,
		MaxAge:     loggerConfig.Keep,
		MaxBackups: loggerConfig.Backups,
		Compress:   loggerConfig.Compress,
	}
}

// NewProductionEncoderConfig
// Load Encoder Config
func NewProductionEncoderConfig() zapcore.EncoderConfig {
	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.TimeKey = "time"
	EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format("2006-01-02 15:04:05"))
	}
	return EncoderConfig
}
