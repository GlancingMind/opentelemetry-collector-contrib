// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

// Deprecated: [v0.108.0] use LeveledMeter instead.
func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/GlancingMind/opentelemetry-collector-contrib/receiver/fluentforwardreceiver")
}

func LeveledMeter(settings component.TelemetrySettings, level configtelemetry.Level) metric.Meter {
	return settings.LeveledMeterProvider(level).Meter("github.com/GlancingMind/opentelemetry-collector-contrib/receiver/fluentforwardreceiver")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/GlancingMind/opentelemetry-collector-contrib/receiver/fluentforwardreceiver")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                   metric.Meter
	FluentClosedConnections metric.Int64UpDownCounter
	FluentEventsParsed      metric.Int64UpDownCounter
	FluentOpenedConnections metric.Int64UpDownCounter
	FluentParseFailures     metric.Int64UpDownCounter
	FluentRecordsGenerated  metric.Int64UpDownCounter
	meters                  map[configtelemetry.Level]metric.Meter
}

// TelemetryBuilderOption applies changes to default builder.
type TelemetryBuilderOption interface {
	apply(*TelemetryBuilder)
}

type telemetryBuilderOptionFunc func(mb *TelemetryBuilder)

func (tbof telemetryBuilderOptionFunc) apply(mb *TelemetryBuilder) {
	tbof(mb)
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
	builder.FluentClosedConnections, err = builder.meters[configtelemetry.LevelBasic].Int64UpDownCounter(
		"otelcol_fluent_closed_connections",
		metric.WithDescription("Number of connections closed to the fluentforward receiver"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.FluentEventsParsed, err = builder.meters[configtelemetry.LevelBasic].Int64UpDownCounter(
		"otelcol_fluent_events_parsed",
		metric.WithDescription("Number of Fluent events parsed successfully"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.FluentOpenedConnections, err = builder.meters[configtelemetry.LevelBasic].Int64UpDownCounter(
		"otelcol_fluent_opened_connections",
		metric.WithDescription("Number of connections opened to the fluentforward receiver"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.FluentParseFailures, err = builder.meters[configtelemetry.LevelBasic].Int64UpDownCounter(
		"otelcol_fluent_parse_failures",
		metric.WithDescription("Number of times Fluent messages failed to be decoded"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.FluentRecordsGenerated, err = builder.meters[configtelemetry.LevelBasic].Int64UpDownCounter(
		"otelcol_fluent_records_generated",
		metric.WithDescription("Number of log records generated from Fluent forward input"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
