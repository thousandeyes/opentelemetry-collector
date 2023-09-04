// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelcol // import "github.com/thousandeyes/opentelemetry-collector/otelcol"

import (
	"github.com/thousandeyes/opentelemetry-collector/component"
	"github.com/thousandeyes/opentelemetry-collector/connector"
	"github.com/thousandeyes/opentelemetry-collector/exporter"
	"github.com/thousandeyes/opentelemetry-collector/extension"
	"github.com/thousandeyes/opentelemetry-collector/processor"
	"github.com/thousandeyes/opentelemetry-collector/receiver"
)

// Factories struct holds in a single type all component factories that
// can be handled by the Config.
type Factories struct {
	// Receivers maps receiver type names in the config to the respective factory.
	Receivers map[component.Type]receiver.Factory

	// Processors maps processor type names in the config to the respective factory.
	Processors map[component.Type]processor.Factory

	// Exporters maps exporter type names in the config to the respective factory.
	Exporters map[component.Type]exporter.Factory

	// Extensions maps extension type names in the config to the respective factory.
	Extensions map[component.Type]extension.Factory

	// Connectors maps connector type names in the config to the respective factory.
	Connectors map[component.Type]connector.Factory
}
