// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/GlancingMind/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/exporter"

	"github.com/GlancingMind/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

func getTelemetryAssets(t require.TestingT) (exporter.Settings, *metadata.TelemetryBuilder) {
	s := setupTestTelemetry()
	st := s.NewSettings()
	ts := st.TelemetrySettings
	tb, err := metadata.NewTelemetryBuilder(ts)
	require.NoError(t, err)
	return st, tb
}
