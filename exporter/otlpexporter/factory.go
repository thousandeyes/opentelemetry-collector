// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpexporter // import "github.com/thousandeyes/opentelemetry-collector/exporter/otlpexporter"

import (
	"context"

	"github.com/thousandeyes/opentelemetry-collector/component"
	"github.com/thousandeyes/opentelemetry-collector/config/configcompression"
	"github.com/thousandeyes/opentelemetry-collector/config/configgrpc"
	"github.com/thousandeyes/opentelemetry-collector/config/configopaque"
	"github.com/thousandeyes/opentelemetry-collector/consumer"
	"github.com/thousandeyes/opentelemetry-collector/exporter"
	"github.com/thousandeyes/opentelemetry-collector/exporter/exporterhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "otlp_lite"
)

// NewFactory creates a factory for OTLP exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, component.StabilityLevelStable),
		exporter.WithMetrics(createMetricsExporter, component.StabilityLevelStable),
		exporter.WithLogs(createLogsExporter, component.StabilityLevelBeta),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		TimeoutSettings: exporterhelper.NewDefaultTimeoutSettings(),
		RetrySettings:   exporterhelper.NewDefaultRetrySettings(),
		QueueSettings:   exporterhelper.NewDefaultQueueSettings(),
		GRPCClientSettings: configgrpc.GRPCClientSettings{
			Headers: map[string]configopaque.String{},
			// Default to gzip compression
			Compression: configcompression.Gzip,
			// We almost read 0 bytes, so no need to tune ReadBufferSize.
			WriteBufferSize: 512 * 1024,
		},
	}
}

func createTracesExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Traces, error) {
	oce, err := newExporter(cfg, set)
	if err != nil {
		return nil, err
	}
	oCfg := cfg.(*Config)
	return exporterhelper.NewTracesExporter(ctx, set, cfg,
		oce.pushTraces,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(oCfg.TimeoutSettings),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithStart(oce.start),
		exporterhelper.WithShutdown(oce.shutdown))
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Metrics, error) {
	oce, err := newExporter(cfg, set)
	if err != nil {
		return nil, err
	}
	oCfg := cfg.(*Config)
	return exporterhelper.NewMetricsExporter(ctx, set, cfg,
		oce.pushMetrics,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(oCfg.TimeoutSettings),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithStart(oce.start),
		exporterhelper.WithShutdown(oce.shutdown),
	)
}

func createLogsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Logs, error) {
	oce, err := newExporter(cfg, set)
	if err != nil {
		return nil, err
	}
	oCfg := cfg.(*Config)
	return exporterhelper.NewLogsExporter(ctx, set, cfg,
		oce.pushLogs,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(oCfg.TimeoutSettings),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithStart(oce.start),
		exporterhelper.WithShutdown(oce.shutdown),
	)
}
