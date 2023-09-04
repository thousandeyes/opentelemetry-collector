module github.com/thousandeyes/opentelemetry-collector

go 1.20

require (
	contrib.go.opencensus.io/exporter/prometheus v0.4.2
	github.com/google/uuid v1.3.1
	github.com/prometheus/client_golang v1.16.0
	github.com/prometheus/client_model v0.4.0
	github.com/prometheus/common v0.44.0
	github.com/shirou/gopsutil/v3 v3.23.7
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.8.4
	go.opencensus.io v0.24.0
	github.com/thousandeyes/opentelemetry-collector/component v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/config/confignet v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/config/configtelemetry v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/confmap v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/connector v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/consumer v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/exporter v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/extension v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/extension/zpagesextension v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/featuregate v1.0.0-rcv0014
	github.com/thousandeyes/opentelemetry-collector/pdata v1.0.0-rcv0014
	github.com/thousandeyes/opentelemetry-collector/processor v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/receiver v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/semconv v0.0.0-20230904160414-bb0c230d9653
	go.opentelemetry.io/contrib/propagators/b3 v1.17.0
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/bridge/opencensus v0.39.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.39.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v0.39.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.16.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.16.0
	go.opentelemetry.io/otel/exporters/prometheus v0.39.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v0.39.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.16.0
	go.opentelemetry.io/otel/metric v1.16.0
	go.opentelemetry.io/otel/sdk v1.16.0
	go.opentelemetry.io/otel/sdk/metric v0.39.0
	go.opentelemetry.io/otel/trace v1.16.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.25.0
	golang.org/x/sys v0.11.0
	gonum.org/v1/gonum v0.14.0
	google.golang.org/grpc v1.57.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/prometheus/statsd_exporter v0.22.7 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	go.opentelemetry.io/contrib/zpages v0.42.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.39.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.16.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	golang.org/x/exp v0.0.0-20230711023510-fffb14384f22 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/thousandeyes/opentelemetry-collector/component => ./component

replace github.com/thousandeyes/opentelemetry-collector/confmap => ./confmap

replace github.com/thousandeyes/opentelemetry-collector/config/confignet => ./config/confignet

replace github.com/thousandeyes/opentelemetry-collector/config/configtelemetry => ./config/configtelemetry

replace github.com/thousandeyes/opentelemetry-collector/connector => ./connector

replace github.com/thousandeyes/opentelemetry-collector/consumer => ./consumer

replace github.com/thousandeyes/opentelemetry-collector/exporter => ./exporter

replace github.com/thousandeyes/opentelemetry-collector/extension => ./extension

replace github.com/thousandeyes/opentelemetry-collector/featuregate => ./featuregate

replace github.com/thousandeyes/opentelemetry-collector/semconv => ./semconv

replace github.com/thousandeyes/opentelemetry-collector/pdata => ./pdata

replace github.com/thousandeyes/opentelemetry-collector/processor => ./processor

replace github.com/thousandeyes/opentelemetry-collector/receiver => ./receiver

replace github.com/thousandeyes/opentelemetry-collector/extension/zpagesextension => ./extension/zpagesextension

retract (
	v0.76.0 // Depends on retracted pdata v1.0.0-rc10 module, use v0.76.1
	v0.69.0 // Release failed, use v0.69.1
	v0.57.1 // Release failed, use v0.57.2
	v0.57.0 // Release failed, use v0.57.2
	v0.32.0 // Contains incomplete metrics transition to proto 0.9.0, random components are not working.
)
