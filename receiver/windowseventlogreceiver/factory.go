// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package windowseventlogreceiver // import "github.com/GlancingMind/opentelemetry-collector-contrib/receiver/windowseventlogreceiver"

import (
	"go.opentelemetry.io/collector/receiver"
)

// NewFactory creates a factory for windowseventlog receiver
func NewFactory() receiver.Factory {
	return newFactoryAdapter()
}
