package providertypeassert

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/helper/schema/providertypeassertexpr"
)

var Analyzer = analysisutils.TypeAssertExprRemovalAnalyzer(
	"providertypeassert",
	providertypeassertexpr.Analyzer,
	schema.PackagePath,
	schema.TypeNameProvider,
)
