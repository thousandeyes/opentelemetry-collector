// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"github.com/thousandeyes/opentelemetry-collector/pdata/internal"
	"github.com/thousandeyes/opentelemetry-collector/pdata/internal/data"
	otlpmetrics "github.com/thousandeyes/opentelemetry-collector/pdata/internal/data/protogen/metrics/v1"
	"github.com/thousandeyes/opentelemetry-collector/pdata/pcommon"
)

// Exemplar is a sample input double measurement.
//
// Exemplars also hold information about the environment when the measurement was recorded,
// for example the span and trace ID of the active span when the exemplar was recorded.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewExemplar function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Exemplar struct {
	orig *otlpmetrics.Exemplar
}

func newExemplar(orig *otlpmetrics.Exemplar) Exemplar {
	return Exemplar{orig}
}

// NewExemplar creates a new empty Exemplar.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewExemplar() Exemplar {
	return newExemplar(&otlpmetrics.Exemplar{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms Exemplar) MoveTo(dest Exemplar) {
	*dest.orig = *ms.orig
	*ms.orig = otlpmetrics.Exemplar{}
}

// Timestamp returns the timestamp associated with this Exemplar.
func (ms Exemplar) Timestamp() pcommon.Timestamp {
	return pcommon.Timestamp(ms.orig.TimeUnixNano)
}

// SetTimestamp replaces the timestamp associated with this Exemplar.
func (ms Exemplar) SetTimestamp(v pcommon.Timestamp) {
	ms.orig.TimeUnixNano = uint64(v)
}

// ValueType returns the type of the value for this Exemplar.
// Calling this function on zero-initialized Exemplar will cause a panic.
func (ms Exemplar) ValueType() ExemplarValueType {
	switch ms.orig.Value.(type) {
	case *otlpmetrics.Exemplar_AsDouble:
		return ExemplarValueTypeDouble
	case *otlpmetrics.Exemplar_AsInt:
		return ExemplarValueTypeInt
	}
	return ExemplarValueTypeEmpty
}

// DoubleValue returns the double associated with this Exemplar.
func (ms Exemplar) DoubleValue() float64 {
	return ms.orig.GetAsDouble()
}

// SetDoubleValue replaces the double associated with this Exemplar.
func (ms Exemplar) SetDoubleValue(v float64) {
	ms.orig.Value = &otlpmetrics.Exemplar_AsDouble{
		AsDouble: v,
	}
}

// IntValue returns the int associated with this Exemplar.
func (ms Exemplar) IntValue() int64 {
	return ms.orig.GetAsInt()
}

// SetIntValue replaces the int associated with this Exemplar.
func (ms Exemplar) SetIntValue(v int64) {
	ms.orig.Value = &otlpmetrics.Exemplar_AsInt{
		AsInt: v,
	}
}

// FilteredAttributes returns the FilteredAttributes associated with this Exemplar.
func (ms Exemplar) FilteredAttributes() pcommon.Map {
	return pcommon.Map(internal.NewMap(&ms.orig.FilteredAttributes))
}

// TraceID returns the traceid associated with this Exemplar.
func (ms Exemplar) TraceID() pcommon.TraceID {
	return pcommon.TraceID(ms.orig.TraceId)
}

// SetTraceID replaces the traceid associated with this Exemplar.
func (ms Exemplar) SetTraceID(v pcommon.TraceID) {
	ms.orig.TraceId = data.TraceID(v)
}

// SpanID returns the spanid associated with this Exemplar.
func (ms Exemplar) SpanID() pcommon.SpanID {
	return pcommon.SpanID(ms.orig.SpanId)
}

// SetSpanID replaces the spanid associated with this Exemplar.
func (ms Exemplar) SetSpanID(v pcommon.SpanID) {
	ms.orig.SpanId = data.SpanID(v)
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms Exemplar) CopyTo(dest Exemplar) {
	dest.SetTimestamp(ms.Timestamp())
	switch ms.ValueType() {
	case ExemplarValueTypeDouble:
		dest.SetDoubleValue(ms.DoubleValue())
	case ExemplarValueTypeInt:
		dest.SetIntValue(ms.IntValue())
	}

	ms.FilteredAttributes().CopyTo(dest.FilteredAttributes())
	dest.SetTraceID(ms.TraceID())
	dest.SetSpanID(ms.SpanID())
}
