// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package integrationtest // import "github.com/GlancingMind/opentelemetry-collector-contrib/exporter/elasticsearchexporter/integrationtest"

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func getDebugFlag(t testing.TB) bool {
	raw := os.Getenv("DEBUG")
	if raw == "" {
		return false
	}
	debug, err := strconv.ParseBool(raw)
	require.NoError(t, err, "debug flag parsing failed")
	return debug
}
