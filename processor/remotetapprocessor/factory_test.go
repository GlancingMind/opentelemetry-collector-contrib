// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotetapprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/GlancingMind/opentelemetry-collector-contrib/processor/remotetapprocessor/internal/metadata"
)

func TestNewFactory(t *testing.T) {
	factory := NewFactory()
	assert.EqualValues(t, metadata.Type, factory.Type())
	config := factory.CreateDefaultConfig()
	assert.NotNil(t, config)
}
