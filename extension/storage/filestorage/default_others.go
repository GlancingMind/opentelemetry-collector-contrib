// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package filestorage // import "github.com/GlancingMind/opentelemetry-collector-contrib/extension/storage/filestorage"

func getDefaultDirectory() string {
	return "/var/lib/otelcol/file_storage"
}
