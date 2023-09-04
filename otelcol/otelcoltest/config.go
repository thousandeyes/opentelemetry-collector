// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelcoltest // import "github.com/thousandeyes/opentelemetry-collector/otelcol/otelcoltest"

import (
	"context"

	"github.com/thousandeyes/opentelemetry-collector/confmap"
	"github.com/thousandeyes/opentelemetry-collector/confmap/converter/expandconverter"
	"github.com/thousandeyes/opentelemetry-collector/confmap/provider/envprovider"
	"github.com/thousandeyes/opentelemetry-collector/confmap/provider/fileprovider"
	"github.com/thousandeyes/opentelemetry-collector/confmap/provider/httpprovider"
	"github.com/thousandeyes/opentelemetry-collector/confmap/provider/yamlprovider"
	"github.com/thousandeyes/opentelemetry-collector/otelcol"
)

// LoadConfig loads a config.Config  from file, and does NOT validate the configuration.
func LoadConfig(fileName string, factories otelcol.Factories) (*otelcol.Config, error) {
	// Read yaml config from file
	provider, err := otelcol.NewConfigProvider(otelcol.ConfigProviderSettings{
		ResolverSettings: confmap.ResolverSettings{
			URIs:       []string{fileName},
			Providers:  makeMapProvidersMap(fileprovider.New(), envprovider.New(), yamlprovider.New(), httpprovider.New()),
			Converters: []confmap.Converter{expandconverter.New()},
		},
	})
	if err != nil {
		return nil, err
	}
	return provider.Get(context.Background(), factories)
}

// LoadConfigAndValidate loads a config from the file, and validates the configuration.
func LoadConfigAndValidate(fileName string, factories otelcol.Factories) (*otelcol.Config, error) {
	cfg, err := LoadConfig(fileName, factories)
	if err != nil {
		return nil, err
	}
	return cfg, cfg.Validate()
}

func makeMapProvidersMap(providers ...confmap.Provider) map[string]confmap.Provider {
	ret := make(map[string]confmap.Provider, len(providers))
	for _, provider := range providers {
		ret[provider.Scheme()] = provider
	}
	return ret
}
