package metric

import (
	"context"
	"errors"
	"fmt"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	otelMetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"time"
)

var meter otelMetric.Meter
var funcDurationHistogram otelMetric.Float64Histogram

func Init() error {
	exporter, err := prometheus.New()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to new prometheus: %v", err))
	}

	ctx := context.Background()
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(config.SERVICENAME),
		),
	)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to new resource: %v", err))
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(exporter),
	)

	otel.SetMeterProvider(meterProvider)
	meter = otel.Meter(config.SERVICENAME)
	funcDurationHistogram, err = meter.Float64Histogram(
		"echo_api_duration",
		otelMetric.WithDescription("duration of func call"),
		otelMetric.WithUnit("ms"),
	)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create histogram, error: %v", err))
	}

	return nil
}

func Duration(ctx context.Context, funName string, start time.Time) {
	duration := time.Since(start)
	label := attribute.Key("func").String(funName)
	funcDurationHistogram.Record(ctx, float64(duration.Nanoseconds()/1000000), otelMetric.WithAttributes(label))
}
