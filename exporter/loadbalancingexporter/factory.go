// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate mdatagen metadata.yaml

package loadbalancingexporter // import "github.com/GlancingMind/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"

	"github.com/GlancingMind/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

// NewFactory creates a factory for the exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, metadata.TracesStability),
		exporter.WithLogs(createLogsExporter, metadata.LogsStability),
		exporter.WithMetrics(createMetricsExporter, metadata.MetricsStability),
	)
}

func createDefaultConfig() component.Config {
	otlpFactory := otlpexporter.NewFactory()
	otlpDefaultCfg := otlpFactory.CreateDefaultConfig().(*otlpexporter.Config)
	otlpDefaultCfg.Endpoint = "placeholder:4317"

	return &Config{
		Protocol: Protocol{
			OTLP: *otlpDefaultCfg,
		},
	}
}

func createTracesExporter(_ context.Context, params exporter.Settings, cfg component.Config) (exporter.Traces, error) {
	return newTracesExporter(params, cfg)
}

func createLogsExporter(_ context.Context, params exporter.Settings, cfg component.Config) (exporter.Logs, error) {
	return newLogsExporter(params, cfg)
}

func createMetricsExporter(_ context.Context, params exporter.Settings, cfg component.Config) (exporter.Metrics, error) {
	return newMetricsExporter(params, cfg)
}
