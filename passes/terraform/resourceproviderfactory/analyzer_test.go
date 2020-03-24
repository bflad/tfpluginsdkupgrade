package resourceproviderfactory_test

import (
	"testing"

	"github.com/bflad/tfpluginsdkupgrade/passes/terraform/resourceproviderfactory"
	"github.com/bflad/tfproviderlint/helper/analysisfixtest"
	_ "github.com/hashicorp/terraform-plugin-sdk/helper/schema" // required for vendoring
	_ "github.com/hashicorp/terraform-plugin-sdk/terraform"     // required for vendoring
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, resourceproviderfactory.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysisfixtest.Run(t, testdata, resourceproviderfactory.Analyzer, "a")
}
