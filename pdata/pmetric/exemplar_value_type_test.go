// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetric // import "github.com/thousandeyes/opentelemetry-collector/pdata/pmetric"

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExemplarValueTypeString(t *testing.T) {
	assert.Equal(t, "Empty", ExemplarValueTypeEmpty.String())
	assert.Equal(t, "Int", ExemplarValueTypeInt.String())
	assert.Equal(t, "Double", ExemplarValueTypeDouble.String())
	assert.Equal(t, "", (ExemplarValueTypeDouble + 1).String())
}
