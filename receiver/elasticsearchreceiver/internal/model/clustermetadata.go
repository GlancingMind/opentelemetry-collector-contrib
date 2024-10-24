// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package model // import "github.com/GlancingMind/opentelemetry-collector-contrib/receiver/elasticsearchreceiver/internal/model"

type ClusterMetadataResponse struct {
	ClusterName string `json:"cluster_name"`
	Version     struct {
		Number string `json:"number"`
	} `json:"version"`
}
