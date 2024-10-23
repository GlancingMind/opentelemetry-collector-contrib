// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package sqlserverreceiver // import "github.com/GlancingMind/opentelemetry-collector-contrib/receiver/sqlserverreceiver"

func (cfg *Config) validateInstanceAndComputerName() error {
	return nil
}
