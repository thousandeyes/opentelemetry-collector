// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelcol

import (
	"github.com/thousandeyes/opentelemetry-collector/connector"
	"github.com/thousandeyes/opentelemetry-collector/connector/connectortest"
	"github.com/thousandeyes/opentelemetry-collector/exporter"
	"github.com/thousandeyes/opentelemetry-collector/exporter/exportertest"
	"github.com/thousandeyes/opentelemetry-collector/extension"
	"github.com/thousandeyes/opentelemetry-collector/extension/extensiontest"
	"github.com/thousandeyes/opentelemetry-collector/processor"
	"github.com/thousandeyes/opentelemetry-collector/processor/processortest"
	"github.com/thousandeyes/opentelemetry-collector/receiver"
	"github.com/thousandeyes/opentelemetry-collector/receiver/receivertest"
)

func nopFactories() (Factories, error) {
	var factories Factories
	var err error

	if factories.Connectors, err = connector.MakeFactoryMap(connectortest.NewNopFactory()); err != nil {
		return Factories{}, err
	}

	if factories.Extensions, err = extension.MakeFactoryMap(extensiontest.NewNopFactory()); err != nil {
		return Factories{}, err
	}

	if factories.Receivers, err = receiver.MakeFactoryMap(receivertest.NewNopFactory()); err != nil {
		return Factories{}, err
	}

	if factories.Exporters, err = exporter.MakeFactoryMap(exportertest.NewNopFactory()); err != nil {
		return Factories{}, err
	}

	if factories.Processors, err = processor.MakeFactoryMap(processortest.NewNopFactory()); err != nil {
		return Factories{}, err
	}

	return factories, err
}
