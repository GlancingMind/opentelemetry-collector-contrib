// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package trace // import "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/trace"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "trace_parser"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new trace parser config with default values
func NewConfig() *Config {
	return NewConfigWithID(operatorType)
}

// NewConfigWithID creates a new trace parser config with default values
func NewConfigWithID(operatorID string) *Config {
	return &Config{
		TransformerConfig: helper.NewTransformerConfig(operatorID, operatorType),
		TraceParser:       helper.NewTraceParser(),
	}
}

// Config is the configuration of a trace parser operator.
type Config struct {
	helper.TransformerConfig `mapstructure:",squash"`
	helper.TraceParser       `mapstructure:",omitempty,squash"`
}

// Build will build a trace parser operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	transformerOperator, err := c.TransformerConfig.Build(set)
	if err != nil {
		return nil, err
	}

	if err := c.TraceParser.Validate(); err != nil {
		return nil, err
	}

	return &Parser{
		TransformerOperator: transformerOperator,
		TraceParser:         c.TraceParser,
	}, nil
}
