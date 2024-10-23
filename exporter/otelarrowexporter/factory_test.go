// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelarrowexporter

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configcompression"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/exportertest"

	"github.com/GlancingMind/opentelemetry-collector-contrib/exporter/otelarrowexporter/internal/arrow"
	"github.com/GlancingMind/opentelemetry-collector-contrib/internal/otelarrow/compression/zstd"
	"github.com/GlancingMind/opentelemetry-collector-contrib/internal/otelarrow/testutil"
)

func TestCreateDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
	ocfg := factory.CreateDefaultConfig().(*Config)
	assert.Equal(t, ocfg.RetryConfig, configretry.NewDefaultBackOffConfig())
	assert.Equal(t, ocfg.QueueSettings, exporterhelper.NewDefaultQueueConfig())
	assert.Equal(t, ocfg.TimeoutSettings, exporterhelper.NewDefaultTimeoutConfig())
	assert.Equal(t, configcompression.TypeZstd, ocfg.Compression)
	assert.Equal(t, ArrowConfig{
		Disabled:           false,
		NumStreams:         max(1, runtime.NumCPU()/2),
		MaxStreamLifetime:  30 * time.Second,
		PayloadCompression: "zstd",
		Zstd:               zstd.DefaultEncoderConfig(),
		Prioritizer:        arrow.DefaultPrioritizer,
	}, ocfg.Arrow)
}

func TestCreateMetrics(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	cfg.ClientConfig.Endpoint = testutil.GetAvailableLocalAddress(t)

	set := exportertest.NewNopSettings()
	oexp, err := factory.CreateMetrics(context.Background(), set, cfg)
	require.NoError(t, err)
	require.NotNil(t, oexp)
}

func TestCreateTraces(t *testing.T) {
	endpoint := testutil.GetAvailableLocalAddress(t)
	tests := []struct {
		name             string
		config           Config
		mustFailOnCreate bool
		mustFailOnStart  bool
	}{
		{
			name: "NoEndpoint",
			config: Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: "",
				},
			},
			mustFailOnCreate: true,
		},
		{
			name: "UseSecure",
			config: Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: endpoint,
					TLSSetting: configtls.ClientConfig{
						Insecure: false,
					},
				},
			},
		},
		{
			name: "Keepalive",
			config: Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: endpoint,
					Keepalive: &configgrpc.KeepaliveClientConfig{
						Time:                30 * time.Second,
						Timeout:             25 * time.Second,
						PermitWithoutStream: true,
					},
				},
			},
		},
		{
			name: "NoneCompression",
			config: Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint:    endpoint,
					Compression: "none",
				},
			},
		},
		{
			name: "GzipCompression",
			config: Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint:    endpoint,
					Compression: configcompression.TypeGzip,
				},
			},
		},
		{
			name: "SnappyCompression",
			config: Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint:    endpoint,
					Compression: configcompression.TypeSnappy,
				},
			},
		},
		{
			name: "ZstdCompression",
			config: Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint:    endpoint,
					Compression: configcompression.TypeZstd,
				},
			},
		},
		{
			name: "Headers",
			config: Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: endpoint,
					Headers: map[string]configopaque.String{
						"hdr1": "val1",
						"hdr2": "val2",
					},
				},
			},
		},
		{
			name: "NumConsumers",
			config: Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: endpoint,
				},
			},
		},
		{
			name: "CaCert",
			config: Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: endpoint,
					TLSSetting: configtls.ClientConfig{
						Config: configtls.Config{
							CAFile: filepath.Join("testdata", "test_cert.pem"),
						},
					},
				},
			},
		},
		{
			name: "CertPemFileError",
			config: Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: endpoint,
					TLSSetting: configtls.ClientConfig{
						Config: configtls.Config{
							CAFile: "nosuchfile",
						},
					},
				},
			},
			mustFailOnStart: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := NewFactory()
			set := exportertest.NewNopSettings()
			cfg := tt.config
			consumer, err := factory.CreateTraces(context.Background(), set, &cfg)
			if tt.mustFailOnCreate {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, consumer)
			err = consumer.Start(context.Background(), componenttest.NewNopHost())
			if tt.mustFailOnStart {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			// Shutdown is called even when Start fails
			err = consumer.Shutdown(context.Background())
			if err != nil {
				// Since the endpoint of OTLP exporter doesn't actually exist,
				// exporter may already stop because it cannot connect.
				assert.Equal(t, "rpc error: code = Canceled desc = grpc: the client connection is closing", err.Error())
			}
		})
	}
}

func TestCreateLogs(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	cfg.ClientConfig.Endpoint = testutil.GetAvailableLocalAddress(t)

	set := exportertest.NewNopSettings()
	oexp, err := factory.CreateLogs(context.Background(), set, cfg)
	require.NoError(t, err)
	require.NotNil(t, oexp)
}

func TestCreateArrowTracesExporter(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	cfg.ClientConfig.Endpoint = testutil.GetAvailableLocalAddress(t)
	cfg.Arrow = ArrowConfig{
		NumStreams: 1,
	}
	set := exportertest.NewNopSettings()
	oexp, err := factory.CreateTraces(context.Background(), set, cfg)
	require.NoError(t, err)
	require.NotNil(t, oexp)
}
