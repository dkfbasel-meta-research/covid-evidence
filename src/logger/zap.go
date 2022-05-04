package logger

import (
	"errors"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {

	mode := os.Getenv("LOG_MODE")
	timestamp := os.Getenv("LOG_TIMESTAMP")
	level := os.Getenv("LOG_LEVEL")

	var logConfig zap.Config

	// use different loggers for production and development
	if mode != "production" {
		logConfig = zap.NewDevelopmentConfig()

	} else {
		logConfig = zap.NewProductionConfig()
	}

	switch strings.ToLower(level) {
	case "debug":
		logConfig.Level.SetLevel(zapcore.DebugLevel)

	case "info":
		logConfig.Level.SetLevel(zapcore.InfoLevel)

	case "warn":
		logConfig.Level.SetLevel(zapcore.WarnLevel)

	case "error":
		logConfig.Level.SetLevel(zapcore.ErrorLevel)

	default:
		// log info and above if nothing else is specified
		logConfig.Level.SetLevel(zapcore.InfoLevel)
	}

	if timestamp == "false" {
		logConfig.EncoderConfig.TimeKey = ""
	}

	if mode != "production" {
		// logout stack trace for errors
		logger, _ = logConfig.Build(
			zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zap.ErrorLevel))

	} else {
		// do not print stack trace for production
		logConfig.DisableStacktrace = true
		logger, _ = logConfig.Build(
			zap.AddCallerSkip(1),
		)
	}

}

// Sync will flush the output of the logger
func Sync() {
	logger.Sync() // nolint:errcheck
}

// Disable will turn of all log outputs
func Disable() {
	logger = zap.NewNop()
}

// EnableDebug will enable the logger in develop mode
func EnableDebug() {
	logConfig := zap.NewDevelopmentConfig()
	logConfig.DisableStacktrace = true
	logger, _ = logConfig.Build()
}

// Zap will return an initialized go.uber.org/zap logger
func Zap() *zap.Logger {
	return logger
}

// Debug will log out an debug statement
func Debug(message string, fields ...zapcore.Field) {
	logger.Debug(message, fields...)
}

// Debug will log out an debug statement
func DebugError(message string, err error, fields ...zapcore.Field) {
	fields = adaptErrors(err, fields)
	logger.Debug(message, fields...)
}

// Info will log out an info statement
func Info(message string, fields ...zapcore.Field) {
	fields = adaptErrors(nil, fields)
	logger.Info(message, fields...)
}

// Warn will log out a warning statement
func Warn(message string, fields ...zapcore.Field) {
	fields = adaptErrors(nil, fields)
	logger.Warn(message, fields...)
}

// Fatal will log out a fatal statement
func Fatal(message string, fields ...zapcore.Field) {
	fields = adaptErrors(nil, fields)
	logger.Fatal(message, fields...)
}

// Fatal will log out a fatal statement
func FatalError(message string, err error, fields ...zapcore.Field) {
	fields = adaptErrors(err, fields)
	logger.Fatal(message, fields...)
}

// Error will log out the given error
func Error(message string, err error, fields ...zapcore.Field) {
	fields = adaptErrors(err, fields)
	logger.Error(message, fields...)
}

// NewError will log out the given error and return an error with the message specified
func NewError(message string, err error, fields ...zapcore.Field) error {
	fields = adaptErrors(err, fields)
	logger.Error(message, fields...)
	return errors.New(message)
}

// Err will return an error without verbose error information
func Err(err error) zapcore.Field {
	return zap.String("error", err.Error())
}

// String will return an string field
func String(name string, text string) zapcore.Field {
	return zap.String(name, text)
}

// Any will accept any data for logging
func Any(name string, data interface{}) zapcore.Field {
	return zap.Reflect(name, data)
}

// UserID will return a string field with name user_id
func UserID(userid string) zapcore.Field {
	return zap.String("user_id", userid)
}

// Timepoint will return a date field with name timepoint
func Timepoint(timepoint time.Time) zapcore.Field {
	return zap.Time("timepoint", timepoint)
}
