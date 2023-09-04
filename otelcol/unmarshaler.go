// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelcol // import "github.com/thousandeyes/opentelemetry-collector/otelcol"

import (
	"go.uber.org/zap/zapcore"

	"github.com/thousandeyes/opentelemetry-collector/config/configtelemetry"
	"github.com/thousandeyes/opentelemetry-collector/confmap"
	"github.com/thousandeyes/opentelemetry-collector/connector"
	"github.com/thousandeyes/opentelemetry-collector/exporter"
	"github.com/thousandeyes/opentelemetry-collector/extension"
	"github.com/thousandeyes/opentelemetry-collector/otelcol/internal/configunmarshaler"
	"github.com/thousandeyes/opentelemetry-collector/processor"
	"github.com/thousandeyes/opentelemetry-collector/receiver"
	"github.com/thousandeyes/opentelemetry-collector/service"
	"github.com/thousandeyes/opentelemetry-collector/service/telemetry"
)

type configSettings struct {
	Receivers  *configunmarshaler.Configs[receiver.Factory]  `mapstructure:"receivers"`
	Processors *configunmarshaler.Configs[processor.Factory] `mapstructure:"processors"`
	Exporters  *configunmarshaler.Configs[exporter.Factory]  `mapstructure:"exporters"`
	Connectors *configunmarshaler.Configs[connector.Factory] `mapstructure:"connectors"`
	Extensions *configunmarshaler.Configs[extension.Factory] `mapstructure:"extensions"`
	Service    service.Config                                `mapstructure:"service"`
}

// unmarshal the configSettings from a confmap.Conf.
// After the config is unmarshalled, `Validate()` must be called to validate.
func unmarshal(v *confmap.Conf, factories Factories) (*configSettings, error) {
	// Unmarshal top level sections and validate.
	cfg := &configSettings{
		Receivers:  configunmarshaler.NewConfigs(factories.Receivers),
		Processors: configunmarshaler.NewConfigs(factories.Processors),
		Exporters:  configunmarshaler.NewConfigs(factories.Exporters),
		Connectors: configunmarshaler.NewConfigs(factories.Connectors),
		Extensions: configunmarshaler.NewConfigs(factories.Extensions),
		// TODO: Add a component.ServiceFactory to allow this to be defined by the Service.
		Service: service.Config{
			Telemetry: telemetry.Config{
				Logs: telemetry.LogsConfig{
					Level:       zapcore.InfoLevel,
					Development: false,
					Encoding:    "console",
					Sampling: &telemetry.LogsSamplingConfig{
						Initial:    100,
						Thereafter: 100,
					},
					OutputPaths:       []string{"stderr"},
					ErrorOutputPaths:  []string{"stderr"},
					DisableCaller:     false,
					DisableStacktrace: false,
					InitialFields:     map[string]any(nil),
				},
				Metrics: telemetry.MetricsConfig{
					Level:   configtelemetry.LevelBasic,
					Address: ":8888",
				},
			},
		},
	}

	return cfg, v.Unmarshal(&cfg, confmap.WithErrorUnused())
}
