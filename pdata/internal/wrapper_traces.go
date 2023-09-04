// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/thousandeyes/opentelemetry-collector/pdata/internal"

import (
	otlpcollectortrace "github.com/thousandeyes/opentelemetry-collector/pdata/internal/data/protogen/collector/trace/v1"
	otlptrace "github.com/thousandeyes/opentelemetry-collector/pdata/internal/data/protogen/trace/v1"
)

type Traces struct {
	orig *otlpcollectortrace.ExportTraceServiceRequest
}

func GetOrigTraces(ms Traces) *otlpcollectortrace.ExportTraceServiceRequest {
	return ms.orig
}

func NewTraces(orig *otlpcollectortrace.ExportTraceServiceRequest) Traces {
	return Traces{orig: orig}
}

// TracesToProto internal helper to convert Traces to protobuf representation.
func TracesToProto(l Traces) otlptrace.TracesData {
	return otlptrace.TracesData{
		ResourceSpans: l.orig.ResourceSpans,
	}
}

// TracesFromProto internal helper to convert protobuf representation to Traces.
func TracesFromProto(orig otlptrace.TracesData) Traces {
	return Traces{orig: &otlpcollectortrace.ExportTraceServiceRequest{
		ResourceSpans: orig.ResourceSpans,
	}}
}
