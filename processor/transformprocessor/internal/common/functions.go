// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package common // import "github.com/GlancingMind/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"

import (
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"
)

func ResourceFunctions() map[string]ottl.Factory[ottlresource.TransformContext] {
	return ottlfuncs.StandardFuncs[ottlresource.TransformContext]()
}

func ScopeFunctions() map[string]ottl.Factory[ottlscope.TransformContext] {
	return ottlfuncs.StandardFuncs[ottlscope.TransformContext]()
}
