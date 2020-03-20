package validatejsonstring_test

import (
	"testing"

	"github.com/bflad/tfpluginsdkupgrade/helper/analysisfixtest"
	"github.com/bflad/tfpluginsdkupgrade/passes/helper/validation/validatejsonstring"
	_ "github.com/hashicorp/terraform-plugin-sdk/helper/validation" // required for vendoring
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, validatejsonstring.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysisfixtest.Run(t, testdata, validatejsonstring.Analyzer, "a")
}
