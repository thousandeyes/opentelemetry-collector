// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetricotlp // import "github.com/thousandeyes/opentelemetry-collector/pdata/pmetric/pmetricotlp"

import (
	"bytes"

	"github.com/thousandeyes/opentelemetry-collector/pdata/internal"
	otlpcollectormetrics "github.com/thousandeyes/opentelemetry-collector/pdata/internal/data/protogen/collector/metrics/v1"
	"github.com/thousandeyes/opentelemetry-collector/pdata/internal/json"
	"github.com/thousandeyes/opentelemetry-collector/pdata/pmetric"
)

var jsonUnmarshaler = &pmetric.JSONUnmarshaler{}

// ExportRequest represents the request for gRPC/HTTP client/server.
// It's a wrapper for pmetric.Metrics data.
type ExportRequest struct {
	orig *otlpcollectormetrics.ExportMetricsServiceRequest
}

// NewExportRequest returns an empty ExportRequest.
func NewExportRequest() ExportRequest {
	return ExportRequest{orig: &otlpcollectormetrics.ExportMetricsServiceRequest{}}
}

// NewExportRequestFromMetrics returns a ExportRequest from pmetric.Metrics.
// Because ExportRequest is a wrapper for pmetric.Metrics,
// any changes to the provided Metrics struct will be reflected in the ExportRequest and vice versa.
func NewExportRequestFromMetrics(md pmetric.Metrics) ExportRequest {
	return ExportRequest{orig: internal.GetOrigMetrics(internal.Metrics(md))}
}

// MarshalProto marshals ExportRequest into proto bytes.
func (ms ExportRequest) MarshalProto() ([]byte, error) {
	return ms.orig.Marshal()
}

// UnmarshalProto unmarshalls ExportRequest from proto bytes.
func (ms ExportRequest) UnmarshalProto(data []byte) error {
	return ms.orig.Unmarshal(data)
}

// MarshalJSON marshals ExportRequest into JSON bytes.
func (ms ExportRequest) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if err := json.Marshal(&buf, ms.orig); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalJSON unmarshalls ExportRequest from JSON bytes.
func (ms ExportRequest) UnmarshalJSON(data []byte) error {
	md, err := jsonUnmarshaler.UnmarshalMetrics(data)
	if err != nil {
		return err
	}
	*ms.orig = *internal.GetOrigMetrics(internal.Metrics(md))
	return nil
}

func (ms ExportRequest) Metrics() pmetric.Metrics {
	return pmetric.Metrics(internal.NewMetrics(ms.orig))
}
