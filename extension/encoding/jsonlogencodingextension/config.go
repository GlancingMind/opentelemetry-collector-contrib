// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jsonlogencodingextension // import "github.com/GlancingMind/opentelemetry-collector-contrib/extension/encoding/jsonlogencodingextension"

import "fmt"

type JSONEncodingMode string

const (
	JSONEncodingModeBodyWithInlineAttributes JSONEncodingMode = "body_with_inline_attributes"
	JSONEncodingModeBody                     JSONEncodingMode = "body"
)

type Config struct {
	// Export raw log string instead of log wrapper
	Mode JSONEncodingMode `mapstructure:"mode,omitempty"`
}

func (c *Config) Validate() error {
	switch c.Mode {
	case JSONEncodingModeBodyWithInlineAttributes:
	case JSONEncodingModeBody:
	default:
		return fmt.Errorf("invalid mode %q", c.Mode)
	}
	return nil
}
