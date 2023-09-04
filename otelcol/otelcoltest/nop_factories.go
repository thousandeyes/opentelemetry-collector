// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelcoltest // import "github.com/thousandeyes/opentelemetry-collector/otelcol/otelcoltest"

import (
	"github.com/thousandeyes/opentelemetry-collector/connector"
	"github.com/thousandeyes/opentelemetry-collector/connector/connectortest"
	"github.com/thousandeyes/opentelemetry-collector/exporter"
	"github.com/thousandeyes/opentelemetry-collector/exporter/exportertest"
	"github.com/thousandeyes/opentelemetry-collector/extension"
	"github.com/thousandeyes/opentelemetry-collector/extension/extensiontest"
	"github.com/thousandeyes/opentelemetry-collector/otelcol"
	"github.com/thousandeyes/opentelemetry-collector/processor"
	"github.com/thousandeyes/opentelemetry-collector/processor/processortest"
	"github.com/thousandeyes/opentelemetry-collector/receiver"
	"github.com/thousandeyes/opentelemetry-collector/receiver/receivertest"
)

// NopFactories returns a otelcol.Factories with all nop factories.
func NopFactories() (otelcol.Factories, error) {
	var factories otelcol.Factories
	var err error

	if factories.Extensions, err = extension.MakeFactoryMap(extensiontest.NewNopFactory()); err != nil {
		return otelcol.Factories{}, err
	}

	if factories.Receivers, err = receiver.MakeFactoryMap(receivertest.NewNopFactory()); err != nil {
		return otelcol.Factories{}, err
	}

	if factories.Exporters, err = exporter.MakeFactoryMap(exportertest.NewNopFactory()); err != nil {
		return otelcol.Factories{}, err
	}

	if factories.Processors, err = processor.MakeFactoryMap(processortest.NewNopFactory()); err != nil {
		return otelcol.Factories{}, err
	}

	if factories.Connectors, err = connector.MakeFactoryMap(connectortest.NewNopFactory()); err != nil {
		return otelcol.Factories{}, err
	}

	return factories, err
}
