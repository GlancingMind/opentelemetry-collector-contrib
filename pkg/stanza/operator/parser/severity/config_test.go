// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package severity

import (
	"path/filepath"
	"testing"

	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/operatortest"
)

func TestUnmarshal(t *testing.T) {
	operatortest.ConfigUnmarshalTests{
		DefaultConfig: NewConfig(),
		TestsFile:     filepath.Join(".", "testdata", "config.yaml"),
		Tests: []operatortest.ConfigUnmarshalTest{
			{
				Name:   "default",
				Expect: NewConfig(),
			},
			{
				Name: "on_error_drop",
				Expect: func() *Config {
					cfg := NewConfig()
					cfg.OnError = "drop"
					return cfg
				}(),
			},
			{
				Name: "parse_from_simple",
				Expect: func() *Config {
					cfg := NewConfig()
					from := entry.NewBodyField("from")
					cfg.ParseFrom = &from
					return cfg
				}(),
			},
			{
				Name: "parse_with_preset",
				Expect: func() *Config {
					cfg := NewConfig()
					from := entry.NewBodyField("from")
					cfg.ParseFrom = &from
					cfg.Preset = "http"
					return cfg
				}(),
			},
		},
	}.Run(t)
}
