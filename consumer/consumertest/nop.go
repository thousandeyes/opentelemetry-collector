// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consumertest // import "github.com/thousandeyes/opentelemetry-collector/consumer/consumertest"

import (
	"context"

	"github.com/thousandeyes/opentelemetry-collector/pdata/plog"
	"github.com/thousandeyes/opentelemetry-collector/pdata/pmetric"
	"github.com/thousandeyes/opentelemetry-collector/pdata/ptrace"
)

// NewNop returns a Consumer that just drops all received data and returns no error.
func NewNop() Consumer {
	return &baseConsumer{
		ConsumeTracesFunc:  func(ctx context.Context, td ptrace.Traces) error { return nil },
		ConsumeMetricsFunc: func(ctx context.Context, md pmetric.Metrics) error { return nil },
		ConsumeLogsFunc:    func(ctx context.Context, ld plog.Logs) error { return nil },
	}
}
