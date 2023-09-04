// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package obsreport // import "github.com/thousandeyes/opentelemetry-collector/obsreport"

import (
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	scopeName = "github.com/thousandeyes/opentelemetry-collector/obsreport"

	nameSep = "/"
)

func recordError(span trace.Span, err error) {
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}
}
