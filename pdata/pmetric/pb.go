// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetric // import "github.com/thousandeyes/opentelemetry-collector/pdata/pmetric"

import (
	"github.com/thousandeyes/opentelemetry-collector/pdata/internal"
	otlpmetrics "github.com/thousandeyes/opentelemetry-collector/pdata/internal/data/protogen/metrics/v1"
)

var _ MarshalSizer = (*ProtoMarshaler)(nil)

type ProtoMarshaler struct{}

func (e *ProtoMarshaler) MarshalMetrics(md Metrics) ([]byte, error) {
	pb := internal.MetricsToProto(internal.Metrics(md))
	return pb.Marshal()
}

func (e *ProtoMarshaler) MetricsSize(md Metrics) int {
	pb := internal.MetricsToProto(internal.Metrics(md))
	return pb.Size()
}

type ProtoUnmarshaler struct{}

func (d *ProtoUnmarshaler) UnmarshalMetrics(buf []byte) (Metrics, error) {
	pb := otlpmetrics.MetricsData{}
	err := pb.Unmarshal(buf)
	return Metrics(internal.MetricsFromProto(pb)), err
}
