package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]*schema.Provider // want "deprecated terraform.ResourceProvider"

func Provider() *schema.Provider { // want "deprecated terraform.ResourceProvider"
	return &schema.Provider{}
}

func f() {
	var _ *schema.Provider = Provider() // want "deprecated terraform.ResourceProvider"

	testAccProvider = Provider().(*schema.Provider) // type assertion removal will be handled separately

	testAccProviders = map[string]*schema.Provider{ // want "deprecated terraform.ResourceProvider"
		"aws": testAccProvider,
	}

	_ = terraform.State{} // keep terraform dependency
}
