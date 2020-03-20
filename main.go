package main

import (
	"github.com/bflad/tfpluginsdkupgrade/passes/helper/validation/cidrnetwork"
	"github.com/bflad/tfpluginsdkupgrade/passes/helper/validation/validatejsonstring"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main([]*analysis.Analyzer{
		cidrnetwork.Analyzer,
		validatejsonstring.Analyzer,
	}...)
}
