// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consumertest // import "github.com/thousandeyes/opentelemetry-collector/consumer/consumertest"

import (
	"context"

	"github.com/thousandeyes/opentelemetry-collector/consumer"
	"github.com/thousandeyes/opentelemetry-collector/pdata/plog"
	"github.com/thousandeyes/opentelemetry-collector/pdata/pmetric"
	"github.com/thousandeyes/opentelemetry-collector/pdata/ptrace"
)

// Consumer is a convenience interface that implements all consumer interfaces.
// It has a private function on it to forbid external users from implementing it
// and, as a result, to allow us to add extra functions without breaking
// compatibility.
type Consumer interface {
	// Capabilities to implement the base consumer functionality.
	Capabilities() consumer.Capabilities

	// ConsumeTraces to implement the consumer.Traces.
	ConsumeTraces(context.Context, ptrace.Traces) error

	// ConsumeMetrics to implement the consumer.Metrics.
	ConsumeMetrics(context.Context, pmetric.Metrics) error

	// ConsumeLogs to implement the consumer.Logs.
	ConsumeLogs(context.Context, plog.Logs) error

	unexported()
}

var _ consumer.Logs = (Consumer)(nil)
var _ consumer.Metrics = (Consumer)(nil)
var _ consumer.Traces = (Consumer)(nil)

type nonMutatingConsumer struct{}

// Capabilities returns the base consumer capabilities.
func (bc nonMutatingConsumer) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

type baseConsumer struct {
	nonMutatingConsumer
	consumer.ConsumeTracesFunc
	consumer.ConsumeMetricsFunc
	consumer.ConsumeLogsFunc
}

func (bc baseConsumer) unexported() {}
