// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package adapter // import "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/adapter"

import (
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/output/file" // Register parsers and transformers for stanza-based log receivers
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/output/stdout"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/container"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/csv"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/json"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/jsonarray"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/keyvalue"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/regex"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/scope"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/severity"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/syslog"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/time"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/trace"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/parser/uri"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/add"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/assignkeys"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/copy"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/filter"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/flatten"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/move"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/noop"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/recombine"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/remove"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/retain"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/router"
	_ "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/unquote"
)
