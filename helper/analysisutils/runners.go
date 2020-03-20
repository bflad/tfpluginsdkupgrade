package analysisutils

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"path/filepath"

	"golang.org/x/tools/go/analysis"
)

// DeprecatedWithReplacementSelectorExprRunner returns an Analyzer runner for deprecated *ast.SelectorExpr with replacement
func DeprecatedWithReplacementSelectorExprRunner(analyzerName string, selectorExprAnalyzer *analysis.Analyzer, oldPackagePath, oldSelectorName, newPackagePath, newSelectorName string) func(*analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		selectorExprs := pass.ResultOf[selectorExprAnalyzer].([]*ast.SelectorExpr)

		for _, selectorExpr := range selectorExprs {
			newSelectorExpr := &ast.SelectorExpr{
				Sel: selectorExpr.Sel,
				X:   selectorExpr.X,
			}

			if oldPackagePath != newPackagePath {
				newSelectorExpr.X = &ast.Ident{
					Name: filepath.Base(newPackagePath),
				}
			}

			if oldSelectorName != newSelectorName {
				newSelectorExpr.Sel = &ast.Ident{
					Name: newSelectorName,
				}
			}

			var selectorExprBuf, newSelectorExprBuf bytes.Buffer

			if err := format.Node(&selectorExprBuf, pass.Fset, selectorExpr); err != nil {
				return nil, fmt.Errorf("error formatting original: %s", err)
			}

			if err := format.Node(&newSelectorExprBuf, pass.Fset, newSelectorExpr); err != nil {
				return nil, fmt.Errorf("error formatting new: %s", err)
			}

			pass.Report(analysis.Diagnostic{
				Pos:     selectorExpr.Pos(),
				End:     selectorExpr.End(),
				Message: fmt.Sprintf("%s: deprecated %s must be replaced with %s", analyzerName, selectorExprBuf.String(), newSelectorExprBuf.String()),
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "Replace",
						TextEdits: []analysis.TextEdit{
							{
								Pos:     selectorExpr.Pos(),
								End:     selectorExpr.End(),
								NewText: newSelectorExprBuf.Bytes(),
							},
						},
					},
				},
			})
		}

		return nil, nil
	}
}
