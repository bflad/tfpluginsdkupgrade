package customizedifffunccontext

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/helper/schema/customizedifffuncinfo"
	"golang.org/x/tools/go/analysis"
)

const analyzerName = "customizedifffunccontext"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  fmt.Sprintf("add missing context.Context parameter to schema.CustomizeDiffFunc"),
	Requires: []*analysis.Analyzer{
		customizedifffuncinfo.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	infos := pass.ResultOf[customizedifffuncinfo.Analyzer].([]*schema.CustomizeDiffFuncInfo)

	for _, info := range infos {
		// Existing: func(*schema.ResourceDiff, interface {}) error
		// New:      func(context.Context, *schema.ResourceDiff, interface {}) error
		newParam := &ast.Field{
			Names: []*ast.Ident{
				{
					Name: "ctx",
				},
			},
			Type: &ast.SelectorExpr{
				Sel: &ast.Ident{
					Name: "Context",
				},
				X: &ast.Ident{
					Name: "context",
				},
			},
		}

		info.Type.Params.List = append([]*ast.Field{newParam}, info.Type.Params.List...)

		var newAstBuf bytes.Buffer

		// Only write the FuncType because rewriting loses comments
		if err := format.Node(&newAstBuf, pass.Fset, info.Type); err != nil {
			return nil, fmt.Errorf("error formatting new: %s", err)
		}

		newText := newAstBuf.Bytes()

		// Keep function name for FuncDecl
		if info.AstFuncDecl != nil {
			newText = append(newText[:4], append([]byte(" "+info.AstFuncDecl.Name.Name), newText[4:]...)...)
		}

		pass.Report(analysis.Diagnostic{
			Pos:     info.Type.Pos(),
			End:     info.Type.End(),
			Message: fmt.Sprintf("%s: add missing context.Context parameter to schema.CustomizeDiffFunc", analyzerName),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "Replace",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     info.Type.Pos(),
							End:     info.Type.End(),
							NewText: newText,
						},
					},
				},
			},
		})
	}

	return nil, nil
}
