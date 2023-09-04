module github.com/thousandeyes/opentelemetry-collector/extension/ballastextension

go 1.20

require (
	github.com/stretchr/testify v1.8.4
	github.com/thousandeyes/opentelemetry-collector v0.0.0-20230904162551-7a8e3a837a30
	github.com/thousandeyes/opentelemetry-collector/component v0.0.0-20230904162551-7a8e3a837a30
	github.com/thousandeyes/opentelemetry-collector/confmap v0.0.0-20230904162551-7a8e3a837a30
	github.com/thousandeyes/opentelemetry-collector/extension v0.0.0-20230904162551-7a8e3a837a30
	go.uber.org/zap v1.25.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/shirou/gopsutil/v3 v3.23.7 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	github.com/thousandeyes/opentelemetry-collector/config/configtelemetry v0.0.0-20230904162551-7a8e3a837a30 // indirect
	github.com/thousandeyes/opentelemetry-collector/featuregate v1.0.0-rcv0014 // indirect
	github.com/thousandeyes/opentelemetry-collector/pdata v1.0.0-rcv0014 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/grpc v1.57.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/thousandeyes/opentelemetry-collector => ../../

replace github.com/thousandeyes/opentelemetry-collector/component => ../../component

replace github.com/thousandeyes/opentelemetry-collector/confmap => ../../confmap

replace github.com/thousandeyes/opentelemetry-collector/exporter => ../../exporter

replace github.com/thousandeyes/opentelemetry-collector/extension => ../

replace github.com/thousandeyes/opentelemetry-collector/featuregate => ../../featuregate

replace github.com/thousandeyes/opentelemetry-collector/pdata => ../../pdata

replace github.com/thousandeyes/opentelemetry-collector/receiver => ../../receiver

replace github.com/thousandeyes/opentelemetry-collector/semconv => ../../semconv

replace github.com/thousandeyes/opentelemetry-collector/extension/zpagesextension => ../zpagesextension

replace github.com/thousandeyes/opentelemetry-collector/consumer => ../../consumer

retract (
	v0.76.0 // Depends on retracted pdata v1.0.0-rc10 module, use v0.76.1
	v0.69.0 // Release failed, use v0.69.1
)

replace github.com/thousandeyes/opentelemetry-collector/processor => ../../processor

replace github.com/thousandeyes/opentelemetry-collector/connector => ../../connector

replace github.com/thousandeyes/opentelemetry-collector/config/confignet => ../../config/confignet

replace github.com/thousandeyes/opentelemetry-collector/config/configtelemetry => ../../config/configtelemetry
