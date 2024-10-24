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
	return settings.MeterProvider.Meter("github.com/GlancingMind/opentelemetry-collector-contrib/receiver/kafkareceiver")
}

func LeveledMeter(settings component.TelemetrySettings, level configtelemetry.Level) metric.Meter {
	return settings.LeveledMeterProvider(level).Meter("github.com/GlancingMind/opentelemetry-collector-contrib/receiver/kafkareceiver")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/GlancingMind/opentelemetry-collector-contrib/receiver/kafkareceiver")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                    metric.Meter
	KafkaReceiverCurrentOffset               metric.Int64Gauge
	KafkaReceiverMessages                    metric.Int64Counter
	KafkaReceiverOffsetLag                   metric.Int64Gauge
	KafkaReceiverPartitionClose              metric.Int64Counter
	KafkaReceiverPartitionStart              metric.Int64Counter
	KafkaReceiverUnmarshalFailedLogRecords   metric.Int64Counter
	KafkaReceiverUnmarshalFailedMetricPoints metric.Int64Counter
	KafkaReceiverUnmarshalFailedSpans        metric.Int64Counter
	meters                                   map[configtelemetry.Level]metric.Meter
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
	builder.KafkaReceiverCurrentOffset, err = builder.meters[configtelemetry.LevelBasic].Int64Gauge(
		"otelcol_kafka_receiver_current_offset",
		metric.WithDescription("Current message offset"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverMessages, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_kafka_receiver_messages",
		metric.WithDescription("Number of received messages"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverOffsetLag, err = builder.meters[configtelemetry.LevelBasic].Int64Gauge(
		"otelcol_kafka_receiver_offset_lag",
		metric.WithDescription("Current offset lag"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverPartitionClose, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_kafka_receiver_partition_close",
		metric.WithDescription("Number of finished partitions"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverPartitionStart, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_kafka_receiver_partition_start",
		metric.WithDescription("Number of started partitions"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverUnmarshalFailedLogRecords, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_kafka_receiver_unmarshal_failed_log_records",
		metric.WithDescription("Number of log records failed to be unmarshaled"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverUnmarshalFailedMetricPoints, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_kafka_receiver_unmarshal_failed_metric_points",
		metric.WithDescription("Number of metric points failed to be unmarshaled"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.KafkaReceiverUnmarshalFailedSpans, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_kafka_receiver_unmarshal_failed_spans",
		metric.WithDescription("Number of spans failed to be unmarshaled"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
