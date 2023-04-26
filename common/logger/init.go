// description:
// @author renshiwei
// Date: 2022/10/5 18:05

package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

var zapLog *zap.Logger
var sugarLog *zap.SugaredLogger

func parseConfigLevelEncoder(levelEncoderName string) zapcore.LevelEncoder {
	switch levelEncoderName {
	case "capitalColor":
		return zapcore.CapitalColorLevelEncoder
	case "capital":
		return zapcore.CapitalLevelEncoder
	case "lowercase":
		return zapcore.LowercaseLevelEncoder
	default:
		return zapcore.CapitalLevelEncoder
	}
}

func InitLog(level, logFormat string) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	// 设置日志级别
	atomLevel := zap.NewAtomicLevelAt(switchLogLevel(level))

	zapConfig := zap.Config{
		Level:         atomLevel, // 日志级别
		Development:   false,     // 开发模式，堆栈跟踪
		DisableCaller: true,
		Encoding:      switchLogFormat(logFormat), // 输出格式 console 或 json
		EncoderConfig: encoderConfig,              // 编码器配置
		//InitialFields: map[string]interface{}{"serviceName": "spikeProxy"}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout"}, // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	var err error
	zapLog, err = zapConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("log init err: %+v", err))
	}
	sugarLog = zapLog.Sugar()

	zap.ReplaceGlobals(zapLog)
}

func switchLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "error":
		return zap.ErrorLevel
	case "warn":
		return zap.WarnLevel
	case "info":
		return zap.InfoLevel
	case "debug":
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}
}

func switchLogFormat(logFormat string) string {
	switch strings.ToLower(logFormat) {
	case "console":
		return "console"
	case "json":
		return "json"
	default:
		return "console"
	}
}
