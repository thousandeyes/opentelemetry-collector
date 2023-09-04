// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consumertest

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/thousandeyes/opentelemetry-collector/pdata/plog"
	"github.com/thousandeyes/opentelemetry-collector/pdata/pmetric"
	"github.com/thousandeyes/opentelemetry-collector/pdata/ptrace"
)

func TestErr(t *testing.T) {
	err := errors.New("my error")
	ec := NewErr(err)
	require.NotNil(t, ec)
	assert.NotPanics(t, ec.unexported)
	assert.Equal(t, err, ec.ConsumeLogs(context.Background(), plog.NewLogs()))
	assert.Equal(t, err, ec.ConsumeMetrics(context.Background(), pmetric.NewMetrics()))
	assert.Equal(t, err, ec.ConsumeTraces(context.Background(), ptrace.NewTraces()))
}
