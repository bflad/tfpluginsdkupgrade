package analysisutils

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/types"
	"path/filepath"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// DeprecatedSelectorExprWithReplacementAstNodeRunner returns an Analyzer runner for deprecated *ast.SelectorExpr with replacement AST
// This function covers cases where replacement requires custom AST handling, which can be done prior to this function.
func DeprecatedSelectorExprWithReplacementAstNodeRunner(analyzerName string, selectorExprAnalyzer *analysis.Analyzer, oldPackagePath string, oldSelectorName string, newNode ast.Node) func(*analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		selectorExprs := pass.ResultOf[selectorExprAnalyzer].([]*ast.SelectorExpr)

		for _, selectorExpr := range selectorExprs {
			var selectorExprBuf, newNodeBuf bytes.Buffer

			if err := format.Node(&selectorExprBuf, pass.Fset, selectorExpr); err != nil {
				return nil, fmt.Errorf("error formatting original: %s", err)
			}

			if err := format.Node(&newNodeBuf, pass.Fset, newNode); err != nil {
				return nil, fmt.Errorf("error formatting new: %s", err)
			}

			pass.Report(analysis.Diagnostic{
				Pos:     selectorExpr.Pos(),
				End:     selectorExpr.End(),
				Message: fmt.Sprintf("%s: deprecated %s must be replaced with %s", analyzerName, selectorExprBuf.String(), newNodeBuf.String()),
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "Replace",
						TextEdits: []analysis.TextEdit{
							{
								Pos:     selectorExpr.Pos(),
								End:     selectorExpr.End(),
								NewText: newNodeBuf.Bytes(),
							},
						},
					},
				},
			})
		}

		return nil, nil
	}
}

// DeprecatedWithReplacementPointerSelectorExprRunner returns an Analyzer runner for deprecated *ast.SelectorExpr with replacement
func DeprecatedWithReplacementPointerSelectorExprRunner(analyzerName string, selectorExprAnalyzer *analysis.Analyzer, oldPackagePath, oldSelectorName, newPackagePath, newSelectorName string) func(*analysis.Pass) (interface{}, error) {
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

			newStarExpr := &ast.StarExpr{
				X: newSelectorExpr,
			}

			var selectorExprBuf, newStarExprBuf bytes.Buffer

			if err := format.Node(&selectorExprBuf, pass.Fset, selectorExpr); err != nil {
				return nil, fmt.Errorf("error formatting original: %s", err)
			}

			if err := format.Node(&newStarExprBuf, pass.Fset, newStarExpr); err != nil {
				return nil, fmt.Errorf("error formatting new: %s", err)
			}

			pass.Report(analysis.Diagnostic{
				Pos:     selectorExpr.Pos(),
				End:     selectorExpr.End(),
				Message: fmt.Sprintf("%s: deprecated %s must be replaced with %s", analyzerName, selectorExprBuf.String(), newStarExprBuf.String()),
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "Replace",
						TextEdits: []analysis.TextEdit{
							{
								Pos:     selectorExpr.Pos(),
								End:     selectorExpr.End(),
								NewText: newStarExprBuf.Bytes(),
							},
						},
					},
				},
			})
		}

		return nil, nil
	}
}

// TypeAssertExprRemovalRunner returns an Analyzer runner for removing *ast.TypeAssertExpr
func TypeAssertExprRemovalRunner(analyzerName string, typeAssertExprAnalyzer *analysis.Analyzer) func(*analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		typeAssertExprs := pass.ResultOf[typeAssertExprAnalyzer].([]*ast.TypeAssertExpr)

		for _, typeAssertExpr := range typeAssertExprs {
			var typeAssertExprBuf, xBuf bytes.Buffer

			if err := format.Node(&typeAssertExprBuf, pass.Fset, typeAssertExpr); err != nil {
				return nil, fmt.Errorf("error formatting original: %s", err)
			}

			if err := format.Node(&xBuf, pass.Fset, typeAssertExpr.X); err != nil {
				return nil, fmt.Errorf("error formatting new: %s", err)
			}

			pass.Report(analysis.Diagnostic{
				Pos:     typeAssertExpr.Pos(),
				End:     typeAssertExpr.End(),
				Message: fmt.Sprintf("%s: %s type assertion should be removed", analyzerName, typeAssertExprBuf.String()),
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "Remove",
						TextEdits: []analysis.TextEdit{
							{
								Pos:     typeAssertExpr.Pos(),
								End:     typeAssertExpr.End(),
								NewText: xBuf.Bytes(),
							},
						},
					},
				},
			})
		}

		return nil, nil
	}
}

// TypeAssertExprRunner returns an Analyzer runner for *ast.TypeAssertExpr
func TypeAssertExprRunner(packageFunc func(ast.Expr, *types.Info, string) bool, selectorName string) func(*analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
		nodeFilter := []ast.Node{
			(*ast.TypeAssertExpr)(nil),
		}
		var result []*ast.TypeAssertExpr

		inspect.Preorder(nodeFilter, func(n ast.Node) {
			typeAssertExpr := n.(*ast.TypeAssertExpr)

			if !packageFunc(typeAssertExpr.Type, pass.TypesInfo, selectorName) {
				return
			}

			result = append(result, typeAssertExpr)
		})

		return result, nil
	}
}
