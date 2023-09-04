// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/thousandeyes/opentelemetry-collector/pdata/internal/cmd/pdatagen/internal"

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
)

const header = `// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".`

// AllPackages is a list of all packages that needs to be generated.
var AllPackages = []*Package{
	pcommon,
	plog,
	plogotlp,
	pmetric,
	pmetricotlp,
	ptrace,
	ptraceotlp,
}

// Package is a struct used to generate files.
type Package struct {
	name        string
	path        string
	imports     []string
	testImports []string
	// Can be any of sliceOfPtrs, sliceOfValues, messageValueStruct, or messagePtrStruct
	structs []baseStruct
}

const newLine = "\n"

// GenerateFiles generates files with the configured data structures for this Package.
func (p *Package) GenerateFiles() error {
	for _, s := range p.structs {
		var sb bytes.Buffer
		generateHeader(&sb, p.name)

		// Add imports
		sb.WriteString("import (" + newLine)
		for _, imp := range p.imports {
			if imp != "" {
				sb.WriteString("\t" + imp + newLine)
			} else {
				sb.WriteString(newLine)
			}
		}
		sb.WriteString(")")

		// Write all structs
		sb.WriteString(newLine + newLine)
		s.generateStruct(&sb)
		sb.WriteString(newLine)

		path := filepath.Join("pdata", p.path, "generated_"+strings.ToLower(s.getName())+".go")
		// ignore gosec complain about permissions being `0644`.
		//nolint:gosec
		if err := os.WriteFile(path, sb.Bytes(), 0644); err != nil {
			return err
		}
	}
	return nil
}

// GenerateTestFiles generates files with tests for the configured data structures for this Package.
func (p *Package) GenerateTestFiles() error {
	for _, s := range p.structs {
		var sb bytes.Buffer
		generateHeader(&sb, p.name)

		// Add imports
		sb.WriteString("import (" + newLine)
		for _, imp := range p.testImports {
			if imp != "" {
				sb.WriteString("\t" + imp + newLine)
			} else {
				sb.WriteString(newLine)
			}
		}
		sb.WriteString(")")

		// Write all tests
		sb.WriteString(newLine + newLine)
		s.generateTests(&sb)
		if !usedByOtherDataTypes(p.name) {
			sb.WriteString(newLine + newLine)
			s.generateTestValueHelpers(&sb)
		}

		path := filepath.Join("pdata", p.path, "generated_"+strings.ToLower(s.getName())+"_test.go")
		// ignore gosec complain about permissions being `0644`.
		//nolint:gosec
		if err := os.WriteFile(path, sb.Bytes(), 0644); err != nil {
			return err
		}
	}
	return nil
}

// GenerateInternalFiles generates files with internal pdata structures for this Package.
func (p *Package) GenerateInternalFiles() error {
	if !usedByOtherDataTypes(p.name) {
		return nil
	}

	for _, s := range p.structs {
		var sb bytes.Buffer
		generateHeader(&sb, "internal")

		// Add imports
		sb.WriteString("import (" + newLine)
		for _, imp := range p.imports {
			if imp != "" {
				sb.WriteString("\t" + imp + newLine)
			} else {
				sb.WriteString(newLine)
			}
		}
		sb.WriteString(")")

		// Write all types and funcs
		s.generateInternal(&sb)
		sb.WriteString(newLine)

		// Write all tests generate value
		sb.WriteString(newLine + newLine)
		s.generateTestValueHelpers(&sb)
		sb.WriteString(newLine)

		path := filepath.Join("pdata", "internal", "generated_wrapper_"+strings.ToLower(s.getName())+".go")
		// ignore gosec complain about permissions being `0644`.
		//nolint:gosec
		if err := os.WriteFile(path, sb.Bytes(), 0644); err != nil {
			return err
		}
	}
	return nil
}

func generateHeader(sb *bytes.Buffer, packageName string) {
	sb.WriteString(header)
	sb.WriteString(newLine + newLine)
	sb.WriteString("package " + packageName)
	sb.WriteString(newLine + newLine)
}

// usedByOtherDataTypes defines if the package is used by other data types and orig fields of the package's structs
// need to be accessible from other pdata packages.
func usedByOtherDataTypes(packageName string) bool {
	return packageName == "pcommon"
}
