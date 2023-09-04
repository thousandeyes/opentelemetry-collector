// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testcomponents // import "github.com/thousandeyes/opentelemetry-collector/service/internal/testcomponents"

import (
	"context"

	"github.com/thousandeyes/opentelemetry-collector/component"
)

type componentState struct {
	started bool
	stopped bool
}

func (cs *componentState) Started() bool {
	return cs.started
}

func (cs *componentState) Stopped() bool {
	return cs.stopped
}

func (cs *componentState) Start(_ context.Context, _ component.Host) error {
	cs.started = true
	return nil
}

func (cs *componentState) Shutdown(_ context.Context) error {
	cs.stopped = true
	return nil
}
