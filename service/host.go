// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package service // import "github.com/thousandeyes/opentelemetry-collector/service"

import (
	"github.com/thousandeyes/opentelemetry-collector/component"
	"github.com/thousandeyes/opentelemetry-collector/connector"
	"github.com/thousandeyes/opentelemetry-collector/exporter"
	"github.com/thousandeyes/opentelemetry-collector/extension"
	"github.com/thousandeyes/opentelemetry-collector/processor"
	"github.com/thousandeyes/opentelemetry-collector/receiver"
	"github.com/thousandeyes/opentelemetry-collector/service/extensions"
	"github.com/thousandeyes/opentelemetry-collector/service/internal/graph"
)

var _ component.Host = (*serviceHost)(nil)

type serviceHost struct {
	asyncErrorChannel chan error
	receivers         *receiver.Builder
	processors        *processor.Builder
	exporters         *exporter.Builder
	connectors        *connector.Builder
	extensions        *extension.Builder

	buildInfo component.BuildInfo

	pipelines         *graph.Graph
	serviceExtensions *extensions.Extensions
}

// ReportFatalError is used to report to the host that the receiver encountered
// a fatal error (i.e.: an error that the instance can't recover from) after
// its start function has already returned.
func (host *serviceHost) ReportFatalError(err error) {
	host.asyncErrorChannel <- err
}

func (host *serviceHost) GetFactory(kind component.Kind, componentType component.Type) component.Factory {
	switch kind {
	case component.KindReceiver:
		return host.receivers.Factory(componentType)
	case component.KindProcessor:
		return host.processors.Factory(componentType)
	case component.KindExporter:
		return host.exporters.Factory(componentType)
	case component.KindConnector:
		return host.connectors.Factory(componentType)
	case component.KindExtension:
		return host.extensions.Factory(componentType)
	}
	return nil
}

func (host *serviceHost) GetExtensions() map[component.ID]component.Component {
	return host.serviceExtensions.GetExtensions()
}

// Deprecated: [0.79.0] This function will be removed in the future.
// Several components in the contrib repository use this function so it cannot be removed
// before those cases are removed. In most cases, use of this function can be replaced by a
// connector. See https://github.com/open-telemetry/opentelemetry-collector/issues/7370 and
// https://github.com/open-telemetry/opentelemetry-collector/pull/7390#issuecomment-1483710184
// for additional information.
func (host *serviceHost) GetExporters() map[component.DataType]map[component.ID]component.Component {
	return host.pipelines.GetExporters()
}
