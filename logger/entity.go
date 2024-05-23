package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Init()
	Get() *logrus.Logger
	Trace(ctx context.Context, args ...interface{})
	Tracef(ctx context.Context, format string, args ...interface{})
	Traceln(ctx context.Context, args ...interface{})
	Debug(ctx context.Context, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Debugln(ctx context.Context, args ...interface{})
	Print(ctx context.Context, args ...interface{})
	Printf(ctx context.Context, format string, args ...interface{})
	Println(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Infoln(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Warnln(ctx context.Context, args ...interface{})
	Error(ctx context.Context, err error, args ...interface{})
	Errorf(ctx context.Context, err error, format string, args ...interface{})
	Errorln(ctx context.Context, err error, args ...interface{})
	Fatal(ctx context.Context, err error, args ...interface{})
	Fatalf(ctx context.Context, err error, format string, args ...interface{})
	Fatalln(ctx context.Context, err error, args ...interface{})
	Panic(ctx context.Context, err error, args ...interface{})
	Panicf(ctx context.Context, err error, format string, args ...interface{})
	Panicln(ctx context.Context, err error, args ...interface{})
	GetLevel() logrus.Level
	GetLogrus() *logrus.Logger
}

type logger struct {
	log *logrus.Logger
}

func NewLogger() Logger {
	return &logger{
		log: logrus.New(),
	}
}

var loggerInstance Logger

const (
	TraceIdKey      = "traceID"
	SpanIdKey       = "spanID"
	SpanParentIdKey = "spanParentID"
	CallerFileKey   = "callerFile"
	CallerFuncKey   = "callerFunc"
	CallerLineKey   = "callerLine"

	LogTypePubSub        = "pubsub"
	LogTypeRest          = "rest"
	LogTypeSoap          = "soap"
	LogTypeFieldKey      = "logType"
	IsServerFieldKey     = "isServer"
	IsRequestFieldKey    = "isRequest"
	UrlFieldKey          = "url"
	MethodFieldKey       = "method"
	HeadersFieldKey      = "headers"
	BodyFieldKey         = "body"
	StatusCodeFieldKey   = "statusCode"
	IsSubscriberFieldKey = "isSubscriber"
	TopicIdFieldKey      = "topicId"
	SubscriberIdFieldKey = "subscriberId"
	MessageIdFieldKey    = "messageId"
	MessageStateFieldKey = "messageState"
	MessageDataFieldKey  = "messageData"
	SoapActionFieldKey   = "soapAction"
	ErrorFieldKey        = "error"
)
