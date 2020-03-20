package cidrnetwork

import (
	"github.com/bflad/tfpluginsdkupgrade/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/bflad/tfproviderlint/passes/helper/validation/cidrnetworkselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"cidrnetwork",
	cidrnetworkselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameCIDRNetwork,
	validation.PackagePath,
	validation.FuncNameIsCIDRNetwork,
)
