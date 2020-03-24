package resourceprovider

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/helper/terraformtype/terraform"
	"github.com/bflad/tfproviderlint/passes/terraform/resourceproviderselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementPointerSelectorExprAnalyzer(
	"resourceprovider",
	resourceproviderselectorexpr.Analyzer,
	terraform.PackagePath,
	terraform.TypeNameResourceProvider,
	schema.PackagePath,
	schema.TypeNameProvider,
)
