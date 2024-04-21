package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// newProductionLogOption return a production encoderConfig with IS08601TimeEncoder
func newProductionLogOption() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return encoderConfig
}
