package logger

import (
	"bnqkl/chain-cms/config"
	"bnqkl/chain-cms/helper"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	Writer *zap.SugaredLogger
}

var logger *Logger

func GetLogger() *Logger {
	return logger
}

func NewLogger(writer *zap.SugaredLogger) *Logger {
	return &Logger{
		Writer: writer,
	}
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.Writer.Debug(args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.Writer.Info(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.Writer.Warn(args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.Writer.Error(args...)
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.ErrorLevel
	}
}

func InitLogger() error {
	config := config.GetConfig()
	logConfig := config.Log
	rootPath := helper.GetRootPath()
	// console
	consoleEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "messgge",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)
	logEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "messgge",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	logEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(logEncoderConfig)
	// access log file
	accessDirPath := filepath.Join(rootPath, "logs/access")
	err := os.MkdirAll(accessDirPath, os.ModePerm)
	if err != nil {
		return err
	}
	accessFilePath := filepath.Join(accessDirPath, "access.log")
	accessLumberjack := &lumberjack.Logger{
		Filename:   accessFilePath,
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge,
	}
	defer accessLumberjack.Close()
	// error log file
	errorDirPath := filepath.Join(rootPath, "logs/error")
	err = os.MkdirAll(errorDirPath, os.ModePerm)
	if err != nil {
		return err
	}
	errorFilePath := filepath.Join(errorDirPath, "error.log")
	errorLumberjack := &lumberjack.Logger{
		Filename:   errorFilePath,
		MaxSize:    logConfig.MaxSize,
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge,
	}
	defer errorLumberjack.Close()
	zapLevel := getLogLevel(logConfig.Level)
	teecore := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), zapLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(accessLumberjack), zapLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(errorLumberjack), zap.ErrorLevel),
	)
	_logger := zap.New(teecore, zap.AddCaller())
	defer _logger.Sync()
	logger = NewLogger(_logger.Sugar())
	return nil
}
