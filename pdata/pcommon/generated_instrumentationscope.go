// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pcommon

import (
	"github.com/thousandeyes/opentelemetry-collector/pdata/internal"
	otlpcommon "github.com/thousandeyes/opentelemetry-collector/pdata/internal/data/protogen/common/v1"
)

// InstrumentationScope is a message representing the instrumentation scope information.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewInstrumentationScope function to create new instances.
// Important: zero-initialized instance is not valid for use.
type InstrumentationScope internal.InstrumentationScope

func newInstrumentationScope(orig *otlpcommon.InstrumentationScope) InstrumentationScope {
	return InstrumentationScope(internal.NewInstrumentationScope(orig))
}

// NewInstrumentationScope creates a new empty InstrumentationScope.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewInstrumentationScope() InstrumentationScope {
	return newInstrumentationScope(&otlpcommon.InstrumentationScope{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms InstrumentationScope) MoveTo(dest InstrumentationScope) {
	*dest.getOrig() = *ms.getOrig()
	*ms.getOrig() = otlpcommon.InstrumentationScope{}
}

func (ms InstrumentationScope) getOrig() *otlpcommon.InstrumentationScope {
	return internal.GetOrigInstrumentationScope(internal.InstrumentationScope(ms))
}

// Name returns the name associated with this InstrumentationScope.
func (ms InstrumentationScope) Name() string {
	return ms.getOrig().Name
}

// SetName replaces the name associated with this InstrumentationScope.
func (ms InstrumentationScope) SetName(v string) {
	ms.getOrig().Name = v
}

// Version returns the version associated with this InstrumentationScope.
func (ms InstrumentationScope) Version() string {
	return ms.getOrig().Version
}

// SetVersion replaces the version associated with this InstrumentationScope.
func (ms InstrumentationScope) SetVersion(v string) {
	ms.getOrig().Version = v
}

// Attributes returns the Attributes associated with this InstrumentationScope.
func (ms InstrumentationScope) Attributes() Map {
	return Map(internal.NewMap(&ms.getOrig().Attributes))
}

// DroppedAttributesCount returns the droppedattributescount associated with this InstrumentationScope.
func (ms InstrumentationScope) DroppedAttributesCount() uint32 {
	return ms.getOrig().DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this InstrumentationScope.
func (ms InstrumentationScope) SetDroppedAttributesCount(v uint32) {
	ms.getOrig().DroppedAttributesCount = v
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms InstrumentationScope) CopyTo(dest InstrumentationScope) {
	dest.SetName(ms.Name())
	dest.SetVersion(ms.Version())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetDroppedAttributesCount(ms.DroppedAttributesCount())
}
