module github.com/thousandeyes/opentelemetry-collector/exporter/otlphttpexporter

go 1.20

require (
	github.com/stretchr/testify v1.8.4
	github.com/thousandeyes/opentelemetry-collector/collector v0.0.0-20230904162551-7a8e3a837a30
	github.com/thousandeyes/opentelemetry-collector/component v0.0.0-20230904135538-a1388b18d8d2
	github.com/thousandeyes/opentelemetry-collector/config/configcompression v0.0.0-20230904162551-7a8e3a837a30
	github.com/thousandeyes/opentelemetry-collector/config/confighttp v0.0.0-20230904162551-7a8e3a837a30
	github.com/thousandeyes/opentelemetry-collector/config/configopaque v0.0.0-20230904162551-7a8e3a837a30
	github.com/thousandeyes/opentelemetry-collector/config/configtls v0.0.0-20230904162551-7a8e3a837a30
	github.com/thousandeyes/opentelemetry-collector/confmap v0.0.0-20230904162551-7a8e3a837a30
	github.com/thousandeyes/opentelemetry-collector/consumer v0.0.0-20230904162551-7a8e3a837a30
	github.com/thousandeyes/opentelemetry-collector/exporter v0.0.0-20230904135538-a1388b18d8d2
	github.com/thousandeyes/opentelemetry-collector/pdata v1.0.0-rcv0014
	github.com/thousandeyes/opentelemetry-collector/receiver v0.0.0-20230904162551-7a8e3a837a30
	github.com/thousandeyes/opentelemetry-collector/receiver/otlpreceiver v0.0.0-20230904162551-7a8e3a837a30
	go.uber.org/zap v1.25.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.9.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	github.com/thousandeyes/opentelemetry-collector/config/configauth v0.0.0-20230904162551-7a8e3a837a30 // indirect
	github.com/thousandeyes/opentelemetry-collector/config/configgrpc v0.0.0-20230904162551-7a8e3a837a30 // indirect
	github.com/thousandeyes/opentelemetry-collector/config/confignet v0.0.0-20230904162551-7a8e3a837a30 // indirect
	github.com/thousandeyes/opentelemetry-collector/config/configtelemetry v0.0.0-20230904162551-7a8e3a837a30 // indirect
	github.com/thousandeyes/opentelemetry-collector/config/internal v0.0.0-20230904162551-7a8e3a837a30 // indirect
	github.com/thousandeyes/opentelemetry-collector/extension v0.0.0-20230904162551-7a8e3a837a30 // indirect
	github.com/thousandeyes/opentelemetry-collector/extension/auth v0.0.0-20230904162551-7a8e3a837a30 // indirect
	github.com/thousandeyes/opentelemetry-collector/featuregate v1.0.0-rcv0014 // indirect
	github.com/thousandeyes/opentelemetry-collector/processor v0.0.0-20230904162551-7a8e3a837a30 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.42.1-0.20230612162650-64be7e574a17 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.42.0 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/thousandeyes/opentelemetry-collector => ../../

replace github.com/thousandeyes/opentelemetry-collector/component => ../../component

replace github.com/thousandeyes/opentelemetry-collector/config/configauth => ../../config/configauth

replace github.com/thousandeyes/opentelemetry-collector/config/configcompression => ../../config/configcompression

replace github.com/thousandeyes/opentelemetry-collector/config/configgrpc => ../../config/configgrpc

replace github.com/thousandeyes/opentelemetry-collector/config/confighttp => ../../config/confighttp

replace github.com/thousandeyes/opentelemetry-collector/config/confignet => ../../config/confignet

replace github.com/thousandeyes/opentelemetry-collector/config/configopaque => ../../config/configopaque

replace github.com/thousandeyes/opentelemetry-collector/config/configtelemetry => ../../config/configtelemetry

replace github.com/thousandeyes/opentelemetry-collector/config/configtls => ../../config/configtls

replace github.com/thousandeyes/opentelemetry-collector/config/internal => ../../config/internal

replace github.com/thousandeyes/opentelemetry-collector/confmap => ../../confmap

replace github.com/thousandeyes/opentelemetry-collector/connector => ../../connector

replace github.com/thousandeyes/opentelemetry-collector/exporter => ../

replace github.com/thousandeyes/opentelemetry-collector/extension => ../../extension

replace github.com/thousandeyes/opentelemetry-collector/extension/auth => ../../extension/auth

replace github.com/thousandeyes/opentelemetry-collector/featuregate => ../../featuregate

replace github.com/thousandeyes/opentelemetry-collector/pdata => ../../pdata

replace github.com/thousandeyes/opentelemetry-collector/processor => ../../processor

replace github.com/thousandeyes/opentelemetry-collector/receiver => ../../receiver

replace github.com/thousandeyes/opentelemetry-collector/receiver/otlpreceiver => ../../receiver/otlpreceiver

replace github.com/thousandeyes/opentelemetry-collector/semconv => ../../semconv

replace github.com/thousandeyes/opentelemetry-collector/extension/zpagesextension => ../../extension/zpagesextension

replace github.com/thousandeyes/opentelemetry-collector/consumer => ../../consumer

retract (
	v0.76.0 // Depends on retracted pdata v1.0.0-rc10 module, use v0.76.1
	v0.69.0 // Release failed, use v0.69.1
)

// ambiguous import: found package cloud.google.com/go/compute/metadata in multiple modules
replace cloud.google.com/go => cloud.google.com/go v0.110.2
