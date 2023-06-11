package xzlogger

import (
	"fmt"

	xzcontext "github.com/averyyan/xz/core/context"
	xzservice "github.com/averyyan/xz/core/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Info(service xzservice.Service, msg string) {
	GetDefaultLogger().Info(fmt.Sprintf("【%s-%s】:%s", service.GetName(), service.GetDescription(), msg))
}

func Error(service xzservice.Service, msg string) {
	GetDefaultLogger().Error(fmt.Sprintf("【%s-%s】:%s", service.GetName(), service.GetDescription(), msg))
}

func GetDefaultLogger() *service {
	services := xzcontext.GetServicesByType(LoggerServiceType)
	for _, serv := range services {
		loggerService, ok := serv.(*service)
		if ok && loggerService.isDefault {
			return loggerService
		}
	}
	return nil
}

func NewLoggerStdoutService() *service {
	logger, _ := NewStdoutLoggerConfig().Build()
	return &service{
		serviceName:        "stdout",
		serviceDescription: "默认的终端输出",
		Logger:             logger,
		isDefault:          true,
	}
}

func NewZapStdoutEncoderConfig() *zapcore.EncoderConfig {
	return &zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func NewStdoutLoggerConfig() *zap.Config {
	StdoutEncoderConfig := NewZapStdoutEncoderConfig()
	return &zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       true,
		Encoding:          "console",
		DisableStacktrace: true,
		EncoderConfig:     *StdoutEncoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}
}
