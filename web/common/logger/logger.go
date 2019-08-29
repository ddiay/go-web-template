package logger

import (
	log "github.com/ddiay/go-log"
)

var WebLogger *log.Logger

func Debug(v interface{}, args ...interface{}) {
	WebLogger.Debug(v, args...)
}

func Info(v interface{}, args ...interface{}) {
	WebLogger.Info(v, args...)
}

func Warn(v interface{}, args ...interface{}) {
	WebLogger.Warn(v, args...)
}

func Error(v interface{}, args ...interface{}) {
	WebLogger.Error(v, args...)
}

func Fatal(v interface{}, args ...interface{}) {
	WebLogger.Fatal(v, args...)
}
