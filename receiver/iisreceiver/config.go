// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package iisreceiver // import "github.com/GlancingMind/opentelemetry-collector-contrib/receiver/iisreceiver"

import (
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/GlancingMind/opentelemetry-collector-contrib/receiver/iisreceiver/internal/metadata"
)

// Config defines configuration for simple prometheus receiver.
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
}
