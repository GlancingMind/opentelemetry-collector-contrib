// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlqueryreceiver // import "github.com/GlancingMind/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/GlancingMind/opentelemetry-collector-contrib/internal/sqlquery"
)

type Config struct {
	sqlquery.Config `mapstructure:",squash"`
}

func createDefaultConfig() component.Config {
	cfg := scraperhelper.NewDefaultControllerConfig()
	cfg.CollectionInterval = 10 * time.Second
	return &Config{
		Config: sqlquery.Config{
			ControllerConfig: cfg,
		},
	}
}
