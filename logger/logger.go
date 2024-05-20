package logger

import (
	"context"
	"os"
	"runtime"

	"github.com/alfin87aa/go-common/configs"
	"github.com/alfin87aa/go-common/constants/datetime"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/natefinch/lumberjack.v2"
)

func (l *logger) Init() {
	config := configs.Configs.Log

	l.log = logrus.StandardLogger()
	levelm, err := logrus.ParseLevel(config.Level)
	if err != nil {
		l.log.SetLevel(logrus.InfoLevel)
	} else {
		l.log.SetLevel(levelm)
	}

	l.log.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		ForceQuote:                true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           datetime.ISO8601,
	})

	var locationPath string
	if config.Location != nil {
		locationPath = *config.Location
	} else {
		locationPath = "./logs/app.log"
	}

	os.MkdirAll(locationPath, os.ModePerm)

	l.log.SetOutput(
		&lumberjack.Logger{
			Filename:   locationPath,
			MaxSize:    config.MaxSize,
			MaxAge:     config.MaxAge,
			MaxBackups: config.MaxBackups,
			Compress:   config.Compress,
		},
	)
}

func (l *logger) Get() *logrus.Logger {
	return l.log
}

func addTraceEntries(ctx context.Context, logger logrus.Ext1FieldLogger) logrus.Ext1FieldLogger {
	sc := trace.SpanContextFromContext(ctx)
	newLogger := logger.
		WithField(TraceIdKey, sc.TraceID().String()).
		WithField(SpanIdKey, sc.SpanID().String()).
		WithField(SpanParentIdKey, ctx.Value(SpanParentIdKey))
	return newLogger
}

func addCallerEntries(logger logrus.Ext1FieldLogger) logrus.Ext1FieldLogger {
	if pc, file, line, ok := runtime.Caller(4); ok {
		newLogger := logger.
			WithField(CallerFileKey, file).
			WithField(CallerFuncKey, runtime.FuncForPC(pc).Name()).
			WithField(CallerLineKey, line)
		return newLogger
	}
	return logger
}

// StdEntries Return entries with trace ID entry from span context,
// span ID entry from span context, and
// span parent ID entry from context
func stdEntries(ctx context.Context, logger logrus.Ext1FieldLogger) logrus.Ext1FieldLogger {
	logger = addTraceEntries(ctx, logger)
	logger = addCallerEntries(logger)
	return logger
}

func (l *logger) Trace(ctx context.Context, args ...interface{}) {
	stdEntries(ctx, l.log).Trace(args...)
}

func (l *logger) Tracef(ctx context.Context, format string, args ...interface{}) {
	stdEntries(ctx, l.log).Tracef(format, args...)
}

func (l *logger) Traceln(ctx context.Context, args ...interface{}) {
	stdEntries(ctx, l.log).Traceln(args...)
}

func (l *logger) Debug(ctx context.Context, args ...interface{}) {
	stdEntries(ctx, l.log).Debug(args...)
}

func (l *logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	stdEntries(ctx, l.log).Debugf(format, args...)
}

func (l *logger) Debugln(ctx context.Context, args ...interface{}) {
	stdEntries(ctx, l.log).Debugln(args...)
}

func (l *logger) Print(ctx context.Context, args ...interface{}) {
	stdEntries(ctx, l.log).Print(args...)
}

func (l *logger) Printf(ctx context.Context, format string, args ...interface{}) {
	stdEntries(ctx, l.log).Printf(format, args...)
}

func (l *logger) Println(ctx context.Context, args ...interface{}) {
	stdEntries(ctx, l.log).Println(args...)
}

func (l *logger) Info(ctx context.Context, args ...interface{}) {
	stdEntries(ctx, l.log).Info(args...)
}

func (l *logger) Infof(ctx context.Context, format string, args ...interface{}) {
	stdEntries(ctx, l.log).Infof(format, args...)
}

func (l *logger) Infoln(ctx context.Context, args ...interface{}) {
	stdEntries(ctx, l.log).Infoln(args...)
}

func (l *logger) Warn(ctx context.Context, args ...interface{}) {
	stdEntries(ctx, l.log).Warn(args...)
}

func (l *logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	stdEntries(ctx, l.log).Warnf(format, args...)
}

func (l *logger) Warnln(ctx context.Context, args ...interface{}) {
	stdEntries(ctx, l.log).Warnln(args...)
}

func (l *logger) Error(ctx context.Context, err error, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.RecordError(err)
	}
	stdEntries(ctx, l.log).WithError(err).Error(args...)
}

func (l *logger) Errorf(ctx context.Context, err error, format string, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.RecordError(err)
	}
	stdEntries(ctx, l.log).WithError(err).Errorf(format, args...)
}

func (l *logger) Errorln(ctx context.Context, err error, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.RecordError(err)
	}
	stdEntries(ctx, l.log).WithError(err).Errorln(args...)
}

func (l *logger) Fatal(ctx context.Context, err error, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.RecordError(err)
	}
	stdEntries(ctx, l.log).WithError(err).Fatal(args...)
}

func (l *logger) Fatalf(ctx context.Context, err error, format string, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.RecordError(err)
	}
	stdEntries(ctx, l.log).WithError(err).Fatalf(format, args...)
}

func (l *logger) Fatalln(ctx context.Context, err error, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.RecordError(err)
	}
	stdEntries(ctx, l.log).WithError(err).Fatalln(args...)
}

func (l *logger) Panic(ctx context.Context, err error, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.RecordError(err)
	}
	stdEntries(ctx, l.log).WithError(err).Panic(args...)
}

func (l *logger) Panicf(ctx context.Context, err error, format string, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.RecordError(err)
	}
	stdEntries(ctx, l.log).WithError(err).Panicf(format, args...)
}

func (l *logger) Panicln(ctx context.Context, err error, args ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.RecordError(err)
	}
	stdEntries(ctx, l.log).WithError(err).Panicln(args...)
}

func (l *logger) GetLevel() logrus.Level {
	return l.log.GetLevel()
}

func (l *logger) GetLogrus() *logrus.Logger {
	return l.log
}

func Trace(ctx context.Context, args ...interface{}) {
	loggerInstance.Trace(ctx, args...)
}

func Tracef(ctx context.Context, format string, args ...interface{}) {
	loggerInstance.Tracef(ctx, format, args...)
}

func Traceln(ctx context.Context, args ...interface{}) {
	loggerInstance.Traceln(ctx, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	loggerInstance.Debug(ctx, args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	loggerInstance.Debugf(ctx, format, args...)
}

func Debugln(ctx context.Context, args ...interface{}) {
	loggerInstance.Debugln(ctx, args...)
}

func Print(ctx context.Context, args ...interface{}) {
	loggerInstance.Print(ctx, args...)
}

func Printf(ctx context.Context, format string, args ...interface{}) {
	loggerInstance.Printf(ctx, format, args...)
}

func Println(ctx context.Context, args ...interface{}) {
	loggerInstance.Println(ctx, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	loggerInstance.Info(ctx, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	loggerInstance.Infof(ctx, format, args...)
}

func Infoln(ctx context.Context, args ...interface{}) {
	loggerInstance.Infoln(ctx, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	loggerInstance.Warn(ctx, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	loggerInstance.Warnf(ctx, format, args...)
}

func Warnln(ctx context.Context, args ...interface{}) {
	loggerInstance.Warnln(ctx, args...)
}

func Error(ctx context.Context, err error, args ...interface{}) {
	loggerInstance.Error(ctx, err, args...)
}

func Errorf(ctx context.Context, err error, format string, args ...interface{}) {
	loggerInstance.Errorf(ctx, err, format, args...)
}

func Errorln(ctx context.Context, err error, args ...interface{}) {
	loggerInstance.Errorln(ctx, err, args...)
}

func Fatal(ctx context.Context, err error, args ...interface{}) {
	loggerInstance.Fatal(ctx, err, args...)
}

func Fatalf(ctx context.Context, err error, format string, args ...interface{}) {
	loggerInstance.Fatalf(ctx, err, format, args...)
}

func Fatalln(ctx context.Context, err error, args ...interface{}) {
	loggerInstance.Fatalln(ctx, err, args...)
}

func Panic(ctx context.Context, err error, args ...interface{}) {
	loggerInstance.Panic(ctx, err, args...)
}

func Panicf(ctx context.Context, err error, format string, args ...interface{}) {
	loggerInstance.Panicf(ctx, err, format, args...)
}

func Panicln(ctx context.Context, err error, args ...interface{}) {
	loggerInstance.Panicln(ctx, err, args...)
}

func GetLevel() logrus.Level {
	return loggerInstance.GetLevel()
}

func GetLogrus() *logrus.Logger {
	return loggerInstance.GetLogrus()
}
