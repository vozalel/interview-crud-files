package logger

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

const NameTraceUUID = "traceUUID"

type Logger struct {
	logger *logrus.Entry
}

func (l Logger) Debug(args ...interface{}) {
	l.log(logrus.DebugLevel, args...)
}

func (l Logger) Info(args ...interface{}) {
	l.log(logrus.InfoLevel, args...)
}

func (l Logger) Warn(args ...interface{}) {
	l.log(logrus.WarnLevel, args...)
}

func (l Logger) Error(args ...interface{}) {
	l.log(logrus.ErrorLevel, args...)
}

func (l Logger) Fatal(args ...interface{}) {
	l.log(logrus.PanicLevel, args...)
}

func (l Logger) WarningWithTraceUUID(ctx context.Context, args ...interface{}) {
	traceUUID := ctx.Value(NameTraceUUID)
	logger := l
	if traceUUID != nil {
		logger = logger.WithField("traceID", traceUUID.(uuid.UUID).String())
	}
	logger.log(logrus.WarnLevel, args...)
}

func (l Logger) WithField(key string, value interface{}) Logger {
	if l.logger != nil {
		return Logger{logger: l.logger.WithField(key, value)}
	}

	return l
}

func (l Logger) WithFields(fields map[string]interface{}) Logger {
	if l.logger != nil {
		return Logger{logger: l.logger.WithFields(fields)}
	}

	return l
}

func (l Logger) log(level logrus.Level, args ...interface{}) {
	if l.logger != nil {
		l.logger.Log(level, args...)
	}
}
