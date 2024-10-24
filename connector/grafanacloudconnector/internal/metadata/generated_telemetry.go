// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

// Deprecated: [v0.108.0] use LeveledMeter instead.
func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/GlancingMind/opentelemetry-collector-contrib/connector/grafanacloudconnector")
}

func LeveledMeter(settings component.TelemetrySettings, level configtelemetry.Level) metric.Meter {
	return settings.LeveledMeterProvider(level).Meter("github.com/GlancingMind/opentelemetry-collector-contrib/connector/grafanacloudconnector")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/GlancingMind/opentelemetry-collector-contrib/connector/grafanacloudconnector")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                        metric.Meter
	GrafanacloudDatapointCount   metric.Int64Counter
	GrafanacloudFlushCount       metric.Int64Counter
	GrafanacloudHostCount        metric.Int64ObservableGauge
	observeGrafanacloudHostCount func(context.Context, metric.Observer) error
	meters                       map[configtelemetry.Level]metric.Meter
}

// TelemetryBuilderOption applies changes to default builder.
type TelemetryBuilderOption interface {
	apply(*TelemetryBuilder)
}

type telemetryBuilderOptionFunc func(mb *TelemetryBuilder)

func (tbof telemetryBuilderOptionFunc) apply(mb *TelemetryBuilder) {
	tbof(mb)
}

// WithGrafanacloudHostCountCallback sets callback for observable GrafanacloudHostCount metric.
func WithGrafanacloudHostCountCallback(cb func() int64, opts ...metric.ObserveOption) TelemetryBuilderOption {
	return telemetryBuilderOptionFunc(func(builder *TelemetryBuilder) {
		builder.observeGrafanacloudHostCount = func(_ context.Context, o metric.Observer) error {
			o.ObserveInt64(builder.GrafanacloudHostCount, cb(), opts...)
			return nil
		}
	})
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...TelemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{meters: map[configtelemetry.Level]metric.Meter{}}
	for _, op := range options {
		op.apply(&builder)
	}
	builder.meters[configtelemetry.LevelBasic] = LeveledMeter(settings, configtelemetry.LevelBasic)
	var err, errs error
	builder.GrafanacloudDatapointCount, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_grafanacloud_datapoint_count",
		metric.WithDescription("Number of datapoints sent to Grafana Cloud"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.GrafanacloudFlushCount, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_grafanacloud_flush_count",
		metric.WithDescription("Number of metrics flushes"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.GrafanacloudHostCount, err = builder.meters[configtelemetry.LevelBasic].Int64ObservableGauge(
		"otelcol_grafanacloud_host_count",
		metric.WithDescription("Number of unique hosts"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	_, err = builder.meters[configtelemetry.LevelBasic].RegisterCallback(builder.observeGrafanacloudHostCount, builder.GrafanacloudHostCount)
	errs = errors.Join(errs, err)
	return &builder, errs
}
