package otel

import (
	"github.com/alfin87aa/go-common/configs"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer   trace.Tracer
	meter    metric.Meter
	counters map[string]metric.Int64Counter
	config   *configs.Config
)

type SpanWrapper struct {
	span trace.Span
}
