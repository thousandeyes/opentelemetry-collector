// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batchprocessor // import "github.com/thousandeyes/opentelemetry-collector/processor/batchprocessor"

import (
	"context"
	"time"

	"github.com/thousandeyes/opentelemetry-collector/component"
	"github.com/thousandeyes/opentelemetry-collector/consumer"
	"github.com/thousandeyes/opentelemetry-collector/internal/obsreportconfig"
	"github.com/thousandeyes/opentelemetry-collector/processor"
)

const (
	// The value of "type" key in configuration.
	typeStr = "batch"

	defaultSendBatchSize = uint32(8192)
	defaultTimeout       = 200 * time.Millisecond

	// defaultMetadataCardinalityLimit should be set to the number
	// of metadata configurations the user expects to submit to
	// the collector.
	defaultMetadataCardinalityLimit = 1000
)

// NewFactory returns a new factory for the Batch processor.
func NewFactory() processor.Factory {
	return processor.NewFactory(
		typeStr,
		createDefaultConfig,
		processor.WithTraces(createTraces, component.StabilityLevelStable),
		processor.WithMetrics(createMetrics, component.StabilityLevelStable),
		processor.WithLogs(createLogs, component.StabilityLevelStable))
}

func createDefaultConfig() component.Config {
	return &Config{
		SendBatchSize:            defaultSendBatchSize,
		Timeout:                  defaultTimeout,
		MetadataCardinalityLimit: defaultMetadataCardinalityLimit,
	}
}

func createTraces(
	_ context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	return newBatchTracesProcessor(set, nextConsumer, cfg.(*Config), obsreportconfig.UseOtelForInternalMetricsfeatureGate.IsEnabled())
}

func createMetrics(
	_ context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (processor.Metrics, error) {
	return newBatchMetricsProcessor(set, nextConsumer, cfg.(*Config), obsreportconfig.UseOtelForInternalMetricsfeatureGate.IsEnabled())
}

func createLogs(
	_ context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	return newBatchLogsProcessor(set, nextConsumer, cfg.(*Config), obsreportconfig.UseOtelForInternalMetricsfeatureGate.IsEnabled())
}
