// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zpagesextension // import "github.com/thousandeyes/opentelemetry-collector/extension/zpagesextension"

import (
	"context"

	"github.com/thousandeyes/opentelemetry-collector/component"
	"github.com/thousandeyes/opentelemetry-collector/config/confignet"
	"github.com/thousandeyes/opentelemetry-collector/extension"
)

const (
	// The value of extension "type" in configuration.
	typeStr = "zpages"

	defaultEndpoint = "localhost:55679"
)

// NewFactory creates a factory for Z-Pages extension.
func NewFactory() extension.Factory {
	return extension.NewFactory(typeStr, createDefaultConfig, createExtension, component.StabilityLevelBeta)
}

func createDefaultConfig() component.Config {
	return &Config{
		TCPAddr: confignet.TCPAddr{
			Endpoint: defaultEndpoint,
		},
	}
}

// createExtension creates the extension based on this config.
func createExtension(_ context.Context, set extension.CreateSettings, cfg component.Config) (extension.Extension, error) {
	return newServer(cfg.(*Config), set.TelemetrySettings), nil
}
