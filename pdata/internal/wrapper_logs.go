// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/thousandeyes/opentelemetry-collector/pdata/internal"

import (
	otlpcollectorlog "github.com/thousandeyes/opentelemetry-collector/pdata/internal/data/protogen/collector/logs/v1"
	otlplogs "github.com/thousandeyes/opentelemetry-collector/pdata/internal/data/protogen/logs/v1"
)

type Logs struct {
	orig *otlpcollectorlog.ExportLogsServiceRequest
}

func GetOrigLogs(ms Logs) *otlpcollectorlog.ExportLogsServiceRequest {
	return ms.orig
}

func NewLogs(orig *otlpcollectorlog.ExportLogsServiceRequest) Logs {
	return Logs{orig: orig}
}

// LogsToProto internal helper to convert Logs to protobuf representation.
func LogsToProto(l Logs) otlplogs.LogsData {
	return otlplogs.LogsData{
		ResourceLogs: l.orig.ResourceLogs,
	}
}

// LogsFromProto internal helper to convert protobuf representation to Logs.
func LogsFromProto(orig otlplogs.LogsData) Logs {
	return Logs{orig: &otlpcollectorlog.ExportLogsServiceRequest{
		ResourceLogs: orig.ResourceLogs,
	}}
}
