package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value("log").(*Logger)
	if !ok {
		panic("no logger in context")
	}
	return logger
}

func NewLogger(level string) (*zap.Logger, error) {
	zapLvl := zap.NewAtomicLevel()


	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")

	zapEncoder := zapcore.NewConsoleEncoder(zapConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), zapLvl),
	)

	zapLogger := zap.New(core, zap.AddCaller())

	return zapLogger, nil
}
