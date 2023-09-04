// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package componenttest // import "github.com/thousandeyes/opentelemetry-collector/component/componenttest"

import (
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/thousandeyes/opentelemetry-collector/component"
	"github.com/thousandeyes/opentelemetry-collector/config/configtelemetry"
	"github.com/thousandeyes/opentelemetry-collector/pdata/pcommon"
)

// NewNopTelemetrySettings returns a new nop telemetry settings for Create* functions.
func NewNopTelemetrySettings() component.TelemetrySettings {
	return component.TelemetrySettings{
		Logger:         zap.NewNop(),
		TracerProvider: trace.NewNoopTracerProvider(),
		MeterProvider:  noop.NewMeterProvider(),
		MetricsLevel:   configtelemetry.LevelNone,
		Resource:       pcommon.NewResource(),
	}
}
