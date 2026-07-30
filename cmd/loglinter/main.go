// Package main provides the standalone command-line entry point for the loglinter tool
package main

import (
	"github.com/kazantsev-developer/loglinter/pkg/linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(linter.Analyzer)
}
