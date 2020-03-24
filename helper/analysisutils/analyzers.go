package analysisutils

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

// DeprecatedWithReplacementPointerSelectorExprAnalyzer returns an Analyzer for deprecated *ast.SelectorExpr with replacement
func DeprecatedWithReplacementPointerSelectorExprAnalyzer(analyzerName string, selectorExprAnalyzer *analysis.Analyzer, oldPackagePath, oldSelectorName, newPackagePath, newSelectorName string) *analysis.Analyzer {
	doc := fmt.Sprintf(`check for deprecated %[2]s.%[3]s usage

The %[1]s analyzer reports usage of the deprecated:

%[2]s.%[3]s

That should be replaced with:

*%[4]s.%[5]s
`, analyzerName, oldPackagePath, oldSelectorName, newPackagePath, newSelectorName)

	return &analysis.Analyzer{
		Name: analyzerName,
		Doc:  doc,
		Requires: []*analysis.Analyzer{
			selectorExprAnalyzer,
		},
		Run: DeprecatedWithReplacementPointerSelectorExprRunner(analyzerName, selectorExprAnalyzer, oldPackagePath, oldSelectorName, newPackagePath, newSelectorName),
	}
}

// TypeAssertExprRemovalAnalyzer returns an Analyzer for *ast.TypeAssertExpr
func TypeAssertExprRemovalAnalyzer(analyzerName string, typeAssertExprAnalyzer *analysis.Analyzer, packagePath string, selectorName string) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: analyzerName,
		Doc:  fmt.Sprintf("remove %s.%s type assertions", packagePath, selectorName),
		Requires: []*analysis.Analyzer{
			typeAssertExprAnalyzer,
		},
		Run: TypeAssertExprRemovalRunner(analyzerName, typeAssertExprAnalyzer),
	}
}

// TypeAssertExprAnalyzer returns an Analyzer for *ast.TypeAssertExpr
func TypeAssertExprAnalyzer(analyzerName string, packageFunc func(ast.Expr, *types.Info, string) bool, packagePath string, selectorName string) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: analyzerName,
		Doc:  fmt.Sprintf("find %s.%s type assertions for later passes", packagePath, selectorName),
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
		Run:        TypeAssertExprRunner(packageFunc, selectorName),
		ResultType: reflect.TypeOf([]*ast.TypeAssertExpr{}),
	}
}
