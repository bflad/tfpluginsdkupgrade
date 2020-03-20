package analysisutils

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
)

// DeprecatedWithReplacementSelectorExprAnalyzer returns an Analyzer for deprecated *ast.SelectorExpr with replacement
func DeprecatedWithReplacementSelectorExprAnalyzer(analyzerName string, selectorExprAnalyzer *analysis.Analyzer, oldPackagePath, oldSelectorName, newPackagePath, newSelectorName string) *analysis.Analyzer {
	doc := fmt.Sprintf(`check for deprecated %[2]s.%[3]s usage

The %[1]s analyzer reports usage of the deprecated:

%[2]s.%[3]s

That should be replaced with:

%[4]s.%[5]s
`, analyzerName, oldPackagePath, oldSelectorName, newPackagePath, newSelectorName)

	return &analysis.Analyzer{
		Name: analyzerName,
		Doc:  doc,
		Requires: []*analysis.Analyzer{
			selectorExprAnalyzer,
		},
		Run: DeprecatedWithReplacementSelectorExprRunner(analyzerName, selectorExprAnalyzer, oldPackagePath, oldSelectorName, newPackagePath, newSelectorName),
	}
}
