package validatejsonstring

import (
	"github.com/bflad/tfpluginsdkupgrade/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/bflad/tfproviderlint/passes/helper/validation/validatejsonstringselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"validatejsonstring",
	validatejsonstringselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameValidateJsonString,
	validation.PackagePath,
	validation.FuncNameStringIsJSON,
)
