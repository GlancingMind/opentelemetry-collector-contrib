// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package influxdbreceiver // import "github.com/GlancingMind/opentelemetry-collector-contrib/receiver/influxdbreceiver"

import (
	"go.opentelemetry.io/collector/config/confighttp"
)

// Config defines configuration for the InfluxDB receiver.
type Config struct {
	confighttp.ServerConfig `mapstructure:",squash"`
}
