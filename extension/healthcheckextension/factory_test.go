// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package healthcheckextension

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/extension/extensiontest"

	"github.com/GlancingMind/opentelemetry-collector-contrib/internal/common/testutil"
)

func TestFactory_CreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig()
	assert.Equal(t, &Config{
		ServerConfig: confighttp.ServerConfig{
			Endpoint: "localhost:13133",
		},
		CheckCollectorPipeline: defaultCheckCollectorPipelineSettings(),
		Path:                   "/",
	}, cfg)

	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
	ext, err := createExtension(context.Background(), extensiontest.NewNopSettings(), cfg)
	require.NoError(t, err)
	require.NotNil(t, ext)
}

func TestFactory_Create(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	cfg.Endpoint = testutil.GetAvailableLocalAddress(t)

	ext, err := createExtension(context.Background(), extensiontest.NewNopSettings(), cfg)
	require.NoError(t, err)
	require.NotNil(t, ext)
}
