package logger

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"

	appcontext "github.com/ishanwardhono/transfer-system/pkg/context"
	"github.com/sirupsen/logrus"
)

const (
	timeLayout = "2006-01-02"
	AppName    = "appname"
)

var (
	log *logrus.Logger
)

func Init(level string) {
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	switch level {
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "warning":
		log.SetLevel(logrus.WarnLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "debug", "all":
		log.SetLevel(logrus.DebugLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	log.Out = os.Stdout
}

func getFileAndLine() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		i := strings.Index(file, AppName)
		if i >= 0 {
			file = file[i:]
		}
	}

	return file, line
}

// Info log
func Info(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Info(args...)
}

// Infof log
func Infof(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Infof(format, args...)
}

// Print log
func Print(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Info(args...)
}

// Printf log
func Printf(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Infof(format, args...)
}

// Debug log
func Debug(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Debug(args...)
}

// Debugf log
func Debugf(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Debugf(format, args...)
}

// Warn log
func Warn(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Warn(args...)
}

// Warnf log
func Warnf(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Warnf(format, args...)
}

// Error log
func Error(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Error(args...)
}

// Errorf log
func Errorf(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Errorf(format, args...)
}

// Fatal log
func Fatal(ctx context.Context, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Fatal(args...)
}

// Fatalf log
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	file, line := getFileAndLine()
	ctxVal := appcontext.GetCtxContent(ctx)
	log.WithField("source", fmt.Sprintf("%s:%d", file, line)).WithField("context", ctxVal).Fatalf(format, args...)
}
