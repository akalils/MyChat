package zlog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"LishengChat-go/internal/config"
	"os"
	"path"
	"runtime"
)

var logger *zap.Logger
var logPath string

// 自动调用
func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间格式为ISO8601。
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志encoder还是JSONEncoder，把日志行格式化成JSON格式
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	conf := config.GetConfig()
	// 从配置文件中获取日志路径，并打开或创建日志文件。
	logPath = conf.LogPath
	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)
	// 将文件句柄包装为同步写入器 fileWriteSyncer，确保日志写入文件时的线程安全性。
	fileWriteSyncer := zapcore.AddSync(file)
	// 使用 zapcore.NewTee 将两个核心组件合并为一个，实现日志同时输出到多个目标的功能。
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),
	)
	// 通过 zap.New 创建一个全局 logger 对象，供后续程序使用。
	logger = zap.New(core)
}

func getFileLogWriter() (writeSyncer zapcore.WriteSyncer) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    100, // 单个文件最大100M
		MaxBackups: 60,  // 多于60个日志文件后，清理较旧的日志
		MaxAge:     1,   // 一天一切割
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}

// getCallerInfoForLog 获得调用方的日志信息，包括函数名，文件名，行号
func getCallerInfoForLog() (callerFields []zap.Field) {
	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的调用方的程序计数器（pc）、文件路径、行号
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name() // 获取函数名
	funcName = path.Base(funcName) // Base函数返回路径的最后一个元素，只保留函数名

	// 将函数名、文件路径和行号封装为 zap.Field 切片并返回。
	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))
	return
}

func Info(message string, fields ...zap.Field) {
	// 获取调用者信息，并将其附加到fields中。
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Debug(message, fields...)
}
