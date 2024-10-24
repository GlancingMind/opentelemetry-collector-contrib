// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"context"

	guuid "github.com/google/uuid"

	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl"
)

func uuid[K any]() (ottl.ExprFunc[K], error) {
	return func(_ context.Context, _ K) (any, error) {
		u := guuid.New()
		return u.String(), nil
	}, nil
}

func createUUIDFunction[K any](_ ottl.FunctionContext, _ ottl.Arguments) (ottl.ExprFunc[K], error) {
	return uuid[K]()
}

func NewUUIDFactory[K any]() ottl.Factory[K] {
	return ottl.NewFactory("UUID", nil, createUUIDFunction[K])
}
