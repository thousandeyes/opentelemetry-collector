// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"

	otlpmetrics "github.com/thousandeyes/opentelemetry-collector/pdata/internal/data/protogen/metrics/v1"
)

func TestSummaryDataPointValueAtQuantileSlice(t *testing.T) {
	es := NewSummaryDataPointValueAtQuantileSlice()
	assert.Equal(t, 0, es.Len())
	es = newSummaryDataPointValueAtQuantileSlice(&[]*otlpmetrics.SummaryDataPoint_ValueAtQuantile{})
	assert.Equal(t, 0, es.Len())

	emptyVal := NewSummaryDataPointValueAtQuantile()
	testVal := generateTestSummaryDataPointValueAtQuantile()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestSummaryDataPointValueAtQuantile(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestSummaryDataPointValueAtQuantileSlice_CopyTo(t *testing.T) {
	dest := NewSummaryDataPointValueAtQuantileSlice()
	// Test CopyTo to empty
	NewSummaryDataPointValueAtQuantileSlice().CopyTo(dest)
	assert.Equal(t, NewSummaryDataPointValueAtQuantileSlice(), dest)

	// Test CopyTo larger slice
	generateTestSummaryDataPointValueAtQuantileSlice().CopyTo(dest)
	assert.Equal(t, generateTestSummaryDataPointValueAtQuantileSlice(), dest)

	// Test CopyTo same size slice
	generateTestSummaryDataPointValueAtQuantileSlice().CopyTo(dest)
	assert.Equal(t, generateTestSummaryDataPointValueAtQuantileSlice(), dest)
}

func TestSummaryDataPointValueAtQuantileSlice_EnsureCapacity(t *testing.T) {
	es := generateTestSummaryDataPointValueAtQuantileSlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestSummaryDataPointValueAtQuantileSlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestSummaryDataPointValueAtQuantileSlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestSummaryDataPointValueAtQuantileSlice(), es)
}

func TestSummaryDataPointValueAtQuantileSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestSummaryDataPointValueAtQuantileSlice()
	dest := NewSummaryDataPointValueAtQuantileSlice()
	src := generateTestSummaryDataPointValueAtQuantileSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestSummaryDataPointValueAtQuantileSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestSummaryDataPointValueAtQuantileSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestSummaryDataPointValueAtQuantileSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestSummaryDataPointValueAtQuantileSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewSummaryDataPointValueAtQuantileSlice()
	emptySlice.RemoveIf(func(el SummaryDataPointValueAtQuantile) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestSummaryDataPointValueAtQuantileSlice()
	pos := 0
	filtered.RemoveIf(func(el SummaryDataPointValueAtQuantile) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestSummaryDataPointValueAtQuantileSlice_Sort(t *testing.T) {
	es := generateTestSummaryDataPointValueAtQuantileSlice()
	es.Sort(func(a, b SummaryDataPointValueAtQuantile) bool {
		return uintptr(unsafe.Pointer(a.orig)) < uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) < uintptr(unsafe.Pointer(es.At(i).orig)))
	}
	es.Sort(func(a, b SummaryDataPointValueAtQuantile) bool {
		return uintptr(unsafe.Pointer(a.orig)) > uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) > uintptr(unsafe.Pointer(es.At(i).orig)))
	}
}

func generateTestSummaryDataPointValueAtQuantileSlice() SummaryDataPointValueAtQuantileSlice {
	es := NewSummaryDataPointValueAtQuantileSlice()
	fillTestSummaryDataPointValueAtQuantileSlice(es)
	return es
}

func fillTestSummaryDataPointValueAtQuantileSlice(es SummaryDataPointValueAtQuantileSlice) {
	*es.orig = make([]*otlpmetrics.SummaryDataPoint_ValueAtQuantile, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlpmetrics.SummaryDataPoint_ValueAtQuantile{}
		fillTestSummaryDataPointValueAtQuantile(newSummaryDataPointValueAtQuantile((*es.orig)[i]))
	}
}
