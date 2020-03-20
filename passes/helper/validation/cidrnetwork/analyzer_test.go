package cidrnetwork_test

import (
	"testing"

	"github.com/bflad/tfpluginsdkupgrade/helper/analysisfixtest"
	"github.com/bflad/tfpluginsdkupgrade/passes/helper/validation/cidrnetwork"
	_ "github.com/hashicorp/terraform-plugin-sdk/helper/validation" // required for vendoring
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, cidrnetwork.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysisfixtest.Run(t, testdata, cidrnetwork.Analyzer, "a")
}
