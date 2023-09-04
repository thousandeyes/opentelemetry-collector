// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configopaque // import "github.com/thousandeyes/opentelemetry-collector/config/configopaque"

import (
	"encoding"
)

// String alias that is marshaled in an opaque way.
type String string

const maskedString = "[REDACTED]"

var _ encoding.TextMarshaler = String("")

// MarshalText marshals the string as `[REDACTED]`.
func (s String) MarshalText() ([]byte, error) {
	return []byte(maskedString), nil
}
