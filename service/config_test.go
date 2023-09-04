// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package service

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	"github.com/thousandeyes/opentelemetry-collector/component"
	"github.com/thousandeyes/opentelemetry-collector/config/configtelemetry"
	"github.com/thousandeyes/opentelemetry-collector/service/extensions"
	"github.com/thousandeyes/opentelemetry-collector/service/pipelines"
	"github.com/thousandeyes/opentelemetry-collector/service/telemetry"
)

func TestConfigValidate(t *testing.T) {
	var testCases = []struct {
		name     string // test case name (also file name containing config yaml)
		cfgFn    func() *Config
		expected error
	}{
		{
			name:     "valid",
			cfgFn:    generateConfig,
			expected: nil,
		},
		{
			name: "custom-service-telemetrySettings-encoding",
			cfgFn: func() *Config {
				cfg := generateConfig()
				cfg.Telemetry.Logs.Encoding = "json"
				return cfg
			},
			expected: nil,
		},
		{
			name: "duplicate-processor-reference",
			cfgFn: func() *Config {
				cfg := generateConfig()
				pipe := cfg.Pipelines[component.NewID("traces")]
				pipe.Processors = append(pipe.Processors, pipe.Processors...)
				return cfg
			},
			expected: fmt.Errorf(`service::pipelines config validation failed: %w`, fmt.Errorf(`pipeline "traces": %w`, errors.New(`references processor "nop" multiple times`))),
		},
		{
			name: "invalid-service-pipeline-type",
			cfgFn: func() *Config {
				cfg := generateConfig()
				cfg.Pipelines[component.NewID("wrongtype")] = &pipelines.PipelineConfig{
					Receivers:  []component.ID{component.NewID("nop")},
					Processors: []component.ID{component.NewID("nop")},
					Exporters:  []component.ID{component.NewID("nop")},
				}
				return cfg
			},
			expected: fmt.Errorf(`service::pipelines config validation failed: %w`, errors.New(`pipeline "wrongtype": unknown datatype "wrongtype"`)),
		},
		{
			name: "invalid-telemetry-metric-config",
			cfgFn: func() *Config {
				cfg := generateConfig()
				cfg.Telemetry.Metrics.Level = configtelemetry.LevelBasic
				cfg.Telemetry.Metrics.Address = ""
				return cfg
			},
			expected: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			cfg := test.cfgFn()
			assert.Equal(t, test.expected, cfg.Validate())
		})
	}
}

func generateConfig() *Config {
	return &Config{
		Telemetry: telemetry.Config{
			Logs: telemetry.LogsConfig{
				Level:             zapcore.DebugLevel,
				Development:       true,
				Encoding:          "console",
				DisableCaller:     true,
				DisableStacktrace: true,
				OutputPaths:       []string{"stderr", "./output-logs"},
				ErrorOutputPaths:  []string{"stderr", "./error-output-logs"},
				InitialFields:     map[string]any{"fieldKey": "filed-value"},
			},
			Metrics: telemetry.MetricsConfig{
				Level:   configtelemetry.LevelNormal,
				Address: ":8080",
			},
		},
		Extensions: extensions.Config{component.NewID("nop")},
		Pipelines: pipelines.Config{
			component.NewID("traces"): {
				Receivers:  []component.ID{component.NewID("nop")},
				Processors: []component.ID{component.NewID("nop")},
				Exporters:  []component.ID{component.NewID("nop")},
			},
		},
	}
}
