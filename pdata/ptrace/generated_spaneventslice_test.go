// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"

	otlptrace "github.com/thousandeyes/opentelemetry-collector/pdata/internal/data/protogen/trace/v1"
)

func TestSpanEventSlice(t *testing.T) {
	es := NewSpanEventSlice()
	assert.Equal(t, 0, es.Len())
	es = newSpanEventSlice(&[]*otlptrace.Span_Event{})
	assert.Equal(t, 0, es.Len())

	emptyVal := NewSpanEvent()
	testVal := generateTestSpanEvent()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestSpanEvent(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestSpanEventSlice_CopyTo(t *testing.T) {
	dest := NewSpanEventSlice()
	// Test CopyTo to empty
	NewSpanEventSlice().CopyTo(dest)
	assert.Equal(t, NewSpanEventSlice(), dest)

	// Test CopyTo larger slice
	generateTestSpanEventSlice().CopyTo(dest)
	assert.Equal(t, generateTestSpanEventSlice(), dest)

	// Test CopyTo same size slice
	generateTestSpanEventSlice().CopyTo(dest)
	assert.Equal(t, generateTestSpanEventSlice(), dest)
}

func TestSpanEventSlice_EnsureCapacity(t *testing.T) {
	es := generateTestSpanEventSlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestSpanEventSlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestSpanEventSlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestSpanEventSlice(), es)
}

func TestSpanEventSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestSpanEventSlice()
	dest := NewSpanEventSlice()
	src := generateTestSpanEventSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestSpanEventSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestSpanEventSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestSpanEventSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestSpanEventSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewSpanEventSlice()
	emptySlice.RemoveIf(func(el SpanEvent) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestSpanEventSlice()
	pos := 0
	filtered.RemoveIf(func(el SpanEvent) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestSpanEventSlice_Sort(t *testing.T) {
	es := generateTestSpanEventSlice()
	es.Sort(func(a, b SpanEvent) bool {
		return uintptr(unsafe.Pointer(a.orig)) < uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) < uintptr(unsafe.Pointer(es.At(i).orig)))
	}
	es.Sort(func(a, b SpanEvent) bool {
		return uintptr(unsafe.Pointer(a.orig)) > uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) > uintptr(unsafe.Pointer(es.At(i).orig)))
	}
}

func generateTestSpanEventSlice() SpanEventSlice {
	es := NewSpanEventSlice()
	fillTestSpanEventSlice(es)
	return es
}

func fillTestSpanEventSlice(es SpanEventSlice) {
	*es.orig = make([]*otlptrace.Span_Event, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlptrace.Span_Event{}
		fillTestSpanEvent(newSpanEvent((*es.orig)[i]))
	}
}
