// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build integration

package mysqlreceiver

import (
	"net"
	"path/filepath"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.opentelemetry.io/collector/component"

	"github.com/GlancingMind/opentelemetry-collector-contrib/internal/coreinternal/scraperinttest"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"
)

const mysqlPort = "3306"

func TestIntegrationWithoutTLS(t *testing.T) {
	scraperinttest.NewIntegrationTest(
		NewFactory(),
		scraperinttest.WithContainerRequest(
			testcontainers.ContainerRequest{
				Image:        "mysql:8.0.33",
				ExposedPorts: []string{mysqlPort},
				WaitingFor: wait.ForListeningPort(mysqlPort).
					WithStartupTimeout(2 * time.Minute),
				Env: map[string]string{
					"MYSQL_ROOT_PASSWORD": "otel",
					"MYSQL_DATABASE":      "otel",
					"MYSQL_USER":          "otel",
					"MYSQL_PASSWORD":      "otel",
				},
				Files: []testcontainers.ContainerFile{
					{
						HostFilePath:      filepath.Join("testdata", "integration", "init.sh"),
						ContainerFilePath: "/docker-entrypoint-initdb.d/init.sh",
						FileMode:          700,
					},
				},
			}),
		scraperinttest.WithCustomConfig(
			func(t *testing.T, cfg component.Config, ci *scraperinttest.ContainerInfo) {
				rCfg := cfg.(*Config)
				rCfg.CollectionInterval = time.Second
				rCfg.Endpoint = net.JoinHostPort(ci.Host(t), ci.MappedPort(t, mysqlPort))
				rCfg.Username = "otel"
				rCfg.Password = "otel"
				// disable TLS connection
				rCfg.TLS.Insecure = true
			}),
		scraperinttest.WithCompareOptions(
			pmetrictest.IgnoreResourceAttributeValue("mysql.instance.endpoint"),
			pmetrictest.IgnoreMetricValues(),
			pmetrictest.IgnoreMetricDataPointsOrder(),
			pmetrictest.IgnoreStartTimestamp(),
			pmetrictest.IgnoreTimestamp(),
		),
	).Run(t)
}

func TestIntegrationWithTLS(t *testing.T) {
	scraperinttest.NewIntegrationTest(
		NewFactory(),
		scraperinttest.WithContainerRequest(
			testcontainers.ContainerRequest{
				Image: "mysql:8.0.33",
				// enable auto TLS certs AND require TLS connections only !
				Cmd:          []string{"--auto_generate_certs=ON", "--require_secure_transport=ON"},
				ExposedPorts: []string{mysqlPort},
				WaitingFor: wait.ForListeningPort(mysqlPort).
					WithStartupTimeout(2 * time.Minute),
				Env: map[string]string{
					"MYSQL_ROOT_PASSWORD": "otel",
					"MYSQL_DATABASE":      "otel",
					"MYSQL_USER":          "otel",
					"MYSQL_PASSWORD":      "otel",
				},
				Files: []testcontainers.ContainerFile{
					{
						HostFilePath:      filepath.Join("testdata", "integration", "init.sh"),
						ContainerFilePath: "/docker-entrypoint-initdb.d/init.sh",
						FileMode:          700,
					},
				},
			}),
		scraperinttest.WithCustomConfig(
			func(t *testing.T, cfg component.Config, ci *scraperinttest.ContainerInfo) {
				rCfg := cfg.(*Config)
				rCfg.CollectionInterval = time.Second
				rCfg.Endpoint = net.JoinHostPort(ci.Host(t), ci.MappedPort(t, mysqlPort))
				rCfg.Username = "otel"
				rCfg.Password = "otel"
				// mysql container is using self-signed certs
				// InsecureSkipVerify will enable TLS but not verify the certificate.
				rCfg.TLS.InsecureSkipVerify = true
			}),
		scraperinttest.WithCompareOptions(
			pmetrictest.IgnoreResourceAttributeValue("mysql.instance.endpoint"),
			pmetrictest.IgnoreMetricValues(),
			pmetrictest.IgnoreMetricDataPointsOrder(),
			pmetrictest.IgnoreStartTimestamp(),
			pmetrictest.IgnoreTimestamp(),
		),
	).Run(t)
}
