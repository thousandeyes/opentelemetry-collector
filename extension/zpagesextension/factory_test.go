// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zpagesextension

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/thousandeyes/opentelemetry-collector/component/componenttest"
	"github.com/thousandeyes/opentelemetry-collector/config/confignet"
	"github.com/thousandeyes/opentelemetry-collector/extension/extensiontest"
	"github.com/thousandeyes/opentelemetry-collector/internal/testutil"
)

func TestFactory_CreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig()
	assert.Equal(t, &Config{
		TCPAddr: confignet.TCPAddr{
			Endpoint: "localhost:55679",
		},
	},
		cfg)

	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
	ext, err := createExtension(context.Background(), extensiontest.NewNopCreateSettings(), cfg)
	require.NoError(t, err)
	require.NotNil(t, ext)
}

func TestFactory_CreateExtension(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	cfg.TCPAddr.Endpoint = testutil.GetAvailableLocalAddress(t)

	ext, err := createExtension(context.Background(), extensiontest.NewNopCreateSettings(), cfg)
	require.NoError(t, err)
	require.NotNil(t, ext)
}
