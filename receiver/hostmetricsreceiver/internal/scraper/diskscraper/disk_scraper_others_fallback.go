// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !linux && !windows

package diskscraper // import "github.com/GlancingMind/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper"

import (
	"github.com/shirou/gopsutil/v4/disk"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

const systemSpecificMetricsLen = 0

func (s *scraper) recordSystemSpecificDataPoints(_ pcommon.Timestamp, _ map[string]disk.IOCountersStat) {
}
