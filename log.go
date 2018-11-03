package kclogs

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

type QYLog struct {
	log *logrus.Logger
	fd  *os.File
}

func (l *QYLog) Debug(format string, args ...interface{}) {
	fields := map[string]interface{}{}
	fields["file"] = fileInfo(2)
	l.DebugWithFields(fields, format, args...)
}

func (l *QYLog) DebugWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if _, ok := fields["file"]; !ok {
		fields["file"] = fileInfo(2)
	}
	l.log.WithFields(fields).Debugf(format, args...)
}

func (l *QYLog) Info(format string, args ...interface{}) {
	fields := map[string]interface{}{}
	fields["file"] = fileInfo(2)
	l.InfoWithFields(fields, format, args...)
}

func (l *QYLog) InfoWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if _, ok := fields["file"]; !ok {
		fields["file"] = fileInfo(2)
	}
	l.log.WithFields(fields).Infof(format, args...)
}

func (l *QYLog) Warn(format string, args ...interface{}) {
	fields := map[string]interface{}{}
	fields["file"] = fileInfo(2)
	l.WarnWithFields(fields, format, args...)
}

func (l *QYLog) WarnWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if _, ok := fields["file"]; !ok {
		fields["file"] = fileInfo(2)
	}
	l.log.WithFields(fields).Warnf(format, args...)
}

func (l *QYLog) Error(format string, args ...interface{}) {
	fields := map[string]interface{}{}
	fields["file"] = fileInfo(2)
	l.ErrorWithFields(fields, format, args...)
}

func (l *QYLog) ErrorWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if _, ok := fields["file"]; !ok {
		fields["file"] = fileInfo(2)
	}
	l.log.WithFields(fields).Errorf(format, args...)
}

func (l *QYLog) Close() {
	if l.fd.Close() != nil {
		fmt.Printf("The fd %v could not closed normally", l.fd)
	}
}

var Log *QYLog

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func InitLog(path, level, format string) {
	logFd, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0666)
	if err != nil {
		fmt.Printf("Cannot open the log file %s with error: %v \n", path, err)
		os.Exit(1)
	}

	Log = &QYLog{log: logrus.New(), fd: logFd}

	switch strings.ToLower(level) {
	case DEBUG:
		Log.log.Level = logrus.DebugLevel
	case ERROR:
		Log.log.Level = logrus.ErrorLevel
	case INFO:
		Log.log.Level = logrus.InfoLevel
	case WARN:
		Log.log.Level = logrus.WarnLevel
	default:
		Log.log.Level = logrus.InfoLevel
	}

	switch strings.ToLower(format) {
	case JSON:
		Log.log.Formatter = &logrus.JSONFormatter{}
	case TEXT:
		Log.log.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: false, DisableSorting: false}
	default:
		Log.log.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: false, DisableSorting: false}
	}

	Log.log.Out = logFd
}
