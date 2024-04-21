package logger

import (
	"flag"
	"fmt"
	"github.com/go-my-rss/go-my-rss/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.SugaredLogger

func init() {
	dev := flag.Bool("dev", false, "run in development mode")
	flag.Parse()

	if *dev {
		development, err := zap.NewDevelopment()
		if err != nil {
			panic(fmt.Errorf("failed to initialize development logger: %w", err))
		}
		logger = development.Named("go-my-rss").Sugar()
	} else {
		logFile, err := os.OpenFile(config.Cfg.Log.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Errorf("failed to open log file: %w", err))
		}
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(newProductionLogOption()),
			zapcore.AddSync(logFile),
			zap.DebugLevel,
		)
		l := zap.New(core)
		logger = l.Named("go-my-rss").Sugar()
	}
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

func Sync() {
	err := logger.Sync()
	if err != nil {
		logger.Errorf("failed to sync logger: %v", err)
	}
}
