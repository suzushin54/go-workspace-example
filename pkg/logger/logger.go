package logger

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
}

func CreateLogger(logLevel *string) *LogrusLogger {
	level, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}

	l := &logrus.Logger{
		Out:   os.Stdout,
		Hooks: nil,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		},
		ReportCaller: false,
		Level:        level,
		ExitFunc:     nil,
		BufferPool:   nil,
	}
	return NewLogrusLogger(l)
}
