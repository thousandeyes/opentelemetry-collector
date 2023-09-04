// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package component // import "github.com/thousandeyes/opentelemetry-collector/component"

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/thousandeyes/opentelemetry-collector/config/configtelemetry"
	"github.com/thousandeyes/opentelemetry-collector/pdata/pcommon"
)

type TelemetrySettings struct {
	// Logger that the factory can use during creation and can pass to the created
	// component to be used later as well.
	Logger *zap.Logger

	// SampledLogger passed to the created component.
	// It will be used for message that can be repeated frequently, to avoid flooding the logs.
	SampledLogger *zap.Logger

	// TracerProvider that the factory can pass to other instrumented third-party libraries.
	TracerProvider trace.TracerProvider

	// MeterProvider that the factory can pass to other instrumented third-party libraries.
	MeterProvider metric.MeterProvider

	// MetricsLevel controls the level of detail for metrics emitted by the collector.
	// Experimental: *NOTE* this field is experimental and may be changed or removed.
	MetricsLevel configtelemetry.Level

	// Resource contains the resource attributes for the collector's telemetry.
	Resource pcommon.Resource
}
