module github.com/thousandeyes/opentelemetry-collector/config/configtls

go 1.20

require (
	github.com/fsnotify/fsnotify v1.6.0
	github.com/stretchr/testify v1.8.4
	github.com/thousandeyes/opentelemetry-collector/config/configopaque v0.0.0-20230904162551-7a8e3a837a30
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/thousandeyes/opentelemetry-collector/config/configopaque => ../configopaque
