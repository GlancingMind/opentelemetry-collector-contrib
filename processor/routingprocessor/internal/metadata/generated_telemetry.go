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
	return settings.MeterProvider.Meter("github.com/GlancingMind/opentelemetry-collector-contrib/processor/routingprocessor")
}

func LeveledMeter(settings component.TelemetrySettings, level configtelemetry.Level) metric.Meter {
	return settings.LeveledMeterProvider(level).Meter("github.com/GlancingMind/opentelemetry-collector-contrib/processor/routingprocessor")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/GlancingMind/opentelemetry-collector-contrib/processor/routingprocessor")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                 metric.Meter
	RoutingProcessorNonRoutedLogRecords   metric.Int64Counter
	RoutingProcessorNonRoutedMetricPoints metric.Int64Counter
	RoutingProcessorNonRoutedSpans        metric.Int64Counter
	meters                                map[configtelemetry.Level]metric.Meter
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
	builder.RoutingProcessorNonRoutedLogRecords, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_routing_processor_non_routed_log_records",
		metric.WithDescription("Number of log records that were not routed to some or all exporters."),
		metric.WithUnit("{records}"),
	)
	errs = errors.Join(errs, err)
	builder.RoutingProcessorNonRoutedMetricPoints, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_routing_processor_non_routed_metric_points",
		metric.WithDescription("Number of metric points that were not routed to some or all exporters."),
		metric.WithUnit("{datapoints}"),
	)
	errs = errors.Join(errs, err)
	builder.RoutingProcessorNonRoutedSpans, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_routing_processor_non_routed_spans",
		metric.WithDescription("Number of spans that were not routed to some or all exporters."),
		metric.WithUnit("{spans}"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
