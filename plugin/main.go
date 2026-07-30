// Package main implements the golangci-lint plugin entry point for loglinter
package main

import (
	"github.com/kazantsev-developer/loglinter/pkg/linter"
	"golang.org/x/tools/go/analysis"
)

type logLinterPlugin struct{}

// NewPluginInstance creates and returns a new plugin configuration instance
func NewPluginInstance() any {
	return &logLinterPlugin{}
}

// GetAnalyzers returns the static analysis tools configured for this plugin execution
func (p *logLinterPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		linter.Analyzer,
	}
}

func main() {}
