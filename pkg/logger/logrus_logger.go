package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	*logrus.Logger
}

func NewLogrusLogger(l *logrus.Logger) *LogrusLogger {
	return &LogrusLogger{
		Logger: l,
	}
}

func (l *LogrusLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.WithContext(ctx).Debugf(format, args...)
}

func (l *LogrusLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.WithContext(ctx).Infof(format, args...)
}

func (l *LogrusLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.WithContext(ctx).Warnf(format, args...)
}

func (l *LogrusLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.WithContext(ctx).Errorf(format, args...)
}

func (l *LogrusLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.WithContext(ctx).Fatalf(format, args...)
}
