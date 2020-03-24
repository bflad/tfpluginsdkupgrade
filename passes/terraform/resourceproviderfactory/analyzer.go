package resourceproviderfactory

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"

	"github.com/bflad/tfpluginsdkupgrade/passes/terraform/resourceproviderfactoryselectorexpr"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"golang.org/x/tools/go/analysis"
)

const analyzerName = "resourceproviderfactory"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  fmt.Sprintf("check for terraform.ResourceProviderFactory that must be replaced"),
	Requires: []*analysis.Analyzer{
		resourceproviderfactoryselectorexpr.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	selectorExprs := pass.ResultOf[resourceproviderfactoryselectorexpr.Analyzer].([]*ast.SelectorExpr)

	for _, selectorExpr := range selectorExprs {
		// func() (*schema.Provider, error)
		newAst := &ast.FuncLit{
			Type: &ast.FuncType{
				Results: &ast.FieldList{
					List: []*ast.Field{
						{
							Type: &ast.StarExpr{
								X: &ast.SelectorExpr{
									Sel: &ast.Ident{
										Name: schema.TypeNameProvider,
									},
									X: &ast.Ident{
										Name: schema.PackageName,
									},
								},
							},
						},
						{
							Type: &ast.Ident{
								Name: "error",
							},
						},
					},
				},
			},
		}

		var newAstBuf, selectorExprBuf bytes.Buffer

		if err := format.Node(&selectorExprBuf, pass.Fset, selectorExpr); err != nil {
			return nil, fmt.Errorf("error formatting original: %s", err)
		}

		if err := format.Node(&newAstBuf, pass.Fset, newAst); err != nil {
			return nil, fmt.Errorf("error formatting new: %s", err)
		}

		pass.Report(analysis.Diagnostic{
			Pos:     selectorExpr.Pos(),
			End:     selectorExpr.End(),
			Message: fmt.Sprintf("%s: %s should be replaced with %s", analyzerName, selectorExprBuf.String(), newAstBuf.String()),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "Replace",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     selectorExpr.Pos(),
							End:     selectorExpr.End(),
							NewText: newAstBuf.Bytes(),
						},
					},
				},
			},
		})
	}

	return nil, nil
}
