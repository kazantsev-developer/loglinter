// Package linter provides a static analysis tool for validating log messages
package linter

import (
	"go/ast"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer checks logging calls for compliance with company-wide rules
var Analyzer = &analysis.Analyzer{
	Name:     "loglinter",
	Doc:      "checks log messages for style, language, and security guidelines",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	getInspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	getInspect.Preorder(nodeFilter, func(n ast.Node) {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}

		selExpr, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		methodName := selExpr.Sel.Name
		if methodName != "Info" && methodName != "Error" && methodName != "Warn" && methodName != "Debug" {
			return
		}

		isLogger := false
		if pass.TypesInfo != nil {
			if tv, ok := pass.TypesInfo.Types[selExpr.X]; ok && tv.Type != nil {
				typeStr := tv.Type.String()
				if stringsContains(typeStr, "log/slog") || stringsContains(typeStr, "zap") {
					isLogger = true
				}
			}
		}

		if !isLogger {
			return
		}

		if len(call.Args) == 0 {
			return
		}

		lit, ok := call.Args[0].(*ast.BasicLit)
		if !ok {
			return
		}

		msg, err := strconv.Unquote(lit.Value)
		if err != nil {
			return
		}

		if errMsg := CheckMessage(msg); errMsg != "" {
			pass.Report(analysis.Diagnostic{
				Pos:     call.Pos(),
				Message: errMsg,
			})
		}
	})

	return nil, nil
}

func stringsContains(s, substr string) bool {
	subLen := len(substr)
	if subLen == 0 {
		return true
	}
	if len(s) < subLen {
		return false
	}

	for i := 0; i <= len(s)-subLen; i++ {
		if s[i:i+subLen] == substr {
			return true
		}
	}
	return false
}
