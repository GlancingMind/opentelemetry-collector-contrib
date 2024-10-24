// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tracking // import "github.com/GlancingMind/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor/internal/tracking"

type MetricPoint struct {
	Identity MetricIdentity
	Value    ValuePoint
}
