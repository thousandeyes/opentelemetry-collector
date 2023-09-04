module github.com/thousandeyes/opentelemetry-collector/config/configauth

go 1.20

require (
	github.com/stretchr/testify v1.8.4
	github.com/thousandeyes/opentelemetry-collector/component v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/extension v0.0.0-20230904160414-bb0c230d9653
	github.com/thousandeyes/opentelemetry-collector/extension/auth v0.0.0-20230904160414-bb0c230d9653
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/thousandeyes/opentelemetry-collector/config/configtelemetry v0.0.0-20230904160414-bb0c230d9653 // indirect
	github.com/thousandeyes/opentelemetry-collector/confmap v0.0.0-20230904160414-bb0c230d9653 // indirect
	github.com/thousandeyes/opentelemetry-collector/featuregate v1.0.0-rcv0014 // indirect
	github.com/thousandeyes/opentelemetry-collector/pdata v1.0.0-rcv0014 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.25.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/grpc v1.57.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/thousandeyes/opentelemetry-collector/pdata => ../../pdata

replace github.com/thousandeyes/opentelemetry-collector/featuregate => ../../featuregate

replace github.com/thousandeyes/opentelemetry-collector/confmap => ../../confmap

replace github.com/thousandeyes/opentelemetry-collector/component => ../../component

replace github.com/thousandeyes/opentelemetry-collector/config/configtelemetry => ../configtelemetry

replace github.com/thousandeyes/opentelemetry-collector/extension => ../../extension

replace github.com/thousandeyes/opentelemetry-collector/extension/auth => ../../extension/auth
