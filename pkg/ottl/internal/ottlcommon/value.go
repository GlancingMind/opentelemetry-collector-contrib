// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlcommon // import "github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl/internal/ottlcommon"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func GetValue(val pcommon.Value) any {
	switch val.Type() {
	case pcommon.ValueTypeStr:
		return val.Str()
	case pcommon.ValueTypeBool:
		return val.Bool()
	case pcommon.ValueTypeInt:
		return val.Int()
	case pcommon.ValueTypeDouble:
		return val.Double()
	case pcommon.ValueTypeMap:
		return val.Map()
	case pcommon.ValueTypeSlice:
		return val.Slice()
	case pcommon.ValueTypeBytes:
		return val.Bytes().AsRaw()
	}
	return nil
}
