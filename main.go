package main

import (
	"github.com/bflad/tfpluginsdkupgrade/passes/helper/schema/customizedifffunccontext"
	"github.com/bflad/tfpluginsdkupgrade/passes/helper/schema/providertypeassert"
	"github.com/bflad/tfpluginsdkupgrade/passes/terraform/resourceprovider"
	"github.com/bflad/tfpluginsdkupgrade/passes/terraform/resourceproviderfactory"
	"github.com/bflad/tfproviderlint/passes/R007"
	"github.com/bflad/tfproviderlint/passes/R008"
	"github.com/bflad/tfproviderlint/passes/V002"
	"github.com/bflad/tfproviderlint/passes/V003"
	"github.com/bflad/tfproviderlint/passes/V004"
	"github.com/bflad/tfproviderlint/passes/V005"
	"github.com/bflad/tfproviderlint/passes/V006"
	"github.com/bflad/tfproviderlint/passes/V007"
	"github.com/bflad/tfproviderlint/passes/V008"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main([]*analysis.Analyzer{
		customizedifffunccontext.Analyzer,
		providertypeassert.Analyzer,
		resourceprovider.Analyzer,
		resourceproviderfactory.Analyzer,
		R007.Analyzer,
		R008.Analyzer,
		V002.Analyzer,
		V003.Analyzer,
		V004.Analyzer,
		V005.Analyzer,
		V006.Analyzer,
		V007.Analyzer,
		V008.Analyzer,
	}...)
}
