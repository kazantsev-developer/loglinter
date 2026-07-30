// Package linter provides a static analysis tool for validating log messages
package linter

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

// TestLogLinter runs integrated static analysis tests using internal testdata suites
func TestLogLinter(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "positive", "negative")
}
