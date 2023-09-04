// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batchprocessor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thousandeyes/opentelemetry-collector/component/componenttest"
	"github.com/thousandeyes/opentelemetry-collector/processor/processortest"
)

func TestCreateDefaultConfig(t *testing.T) {
	factory := NewFactory()

	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
}

func TestCreateProcessor(t *testing.T) {
	factory := NewFactory()

	cfg := factory.CreateDefaultConfig()
	creationSet := processortest.NewNopCreateSettings()
	tp, err := factory.CreateTracesProcessor(context.Background(), creationSet, cfg, nil)
	assert.NotNil(t, tp)
	assert.NoError(t, err, "cannot create trace processor")

	mp, err := factory.CreateMetricsProcessor(context.Background(), creationSet, cfg, nil)
	assert.NotNil(t, mp)
	assert.NoError(t, err, "cannot create metric processor")

	lp, err := factory.CreateLogsProcessor(context.Background(), creationSet, cfg, nil)
	assert.NotNil(t, lp)
	assert.NoError(t, err, "cannot create logs processor")
}
