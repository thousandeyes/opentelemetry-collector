// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpexporter

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/thousandeyes/opentelemetry-collector/component"
	"github.com/thousandeyes/opentelemetry-collector/config/configauth"
	"github.com/thousandeyes/opentelemetry-collector/config/configgrpc"
	"github.com/thousandeyes/opentelemetry-collector/config/configopaque"
	"github.com/thousandeyes/opentelemetry-collector/config/configtls"
	"github.com/thousandeyes/opentelemetry-collector/confmap"
	"github.com/thousandeyes/opentelemetry-collector/confmap/confmaptest"
	"github.com/thousandeyes/opentelemetry-collector/exporter/exporterhelper"
)

func TestUnmarshalDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NoError(t, component.UnmarshalConfig(confmap.New(), cfg))
	assert.Equal(t, factory.CreateDefaultConfig(), cfg)
}

func TestUnmarshalConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NoError(t, component.UnmarshalConfig(cm, cfg))
	assert.Equal(t,
		&Config{
			TimeoutSettings: exporterhelper.TimeoutSettings{
				Timeout: 10 * time.Second,
			},
			RetrySettings: exporterhelper.RetrySettings{
				Enabled:             true,
				InitialInterval:     10 * time.Second,
				RandomizationFactor: 0.7,
				Multiplier:          1.3,
				MaxInterval:         1 * time.Minute,
				MaxElapsedTime:      10 * time.Minute,
			},
			QueueSettings: exporterhelper.QueueSettings{
				Enabled:      true,
				NumConsumers: 2,
				QueueSize:    10,
			},
			GRPCClientSettings: configgrpc.GRPCClientSettings{
				Headers: map[string]configopaque.String{
					"can you have a . here?": "F0000000-0000-0000-0000-000000000000",
					"header1":                "234",
					"another":                "somevalue",
				},
				Endpoint:    "1.2.3.4:1234",
				Compression: "gzip",
				TLSSetting: configtls.TLSClientSetting{
					TLSSetting: configtls.TLSSetting{
						CAFile: "/var/lib/mycert.pem",
					},
					Insecure: false,
				},
				Keepalive: &configgrpc.KeepaliveClientConfig{
					Time:                20 * time.Second,
					PermitWithoutStream: true,
					Timeout:             30 * time.Second,
				},
				WriteBufferSize: 512 * 1024,
				BalancerName:    "round_robin",
				Auth:            &configauth.Authentication{AuthenticatorID: component.NewID("nop")},
			},
		}, cfg)
}
