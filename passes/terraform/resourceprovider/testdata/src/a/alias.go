package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	tf "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviderAlias *schema.Provider
var testAccProvidersAlias map[string]tf.ResourceProvider // want "deprecated tf.ResourceProvider"

func ProviderAlias() tf.ResourceProvider { // want "deprecated tf.ResourceProvider"
	return &schema.Provider{}
}

func fAlias() {
	var _ tf.ResourceProvider = ProviderAlias() // want "deprecated tf.ResourceProvider"

	testAccProviderAlias = ProviderAlias().(*schema.Provider) // type assertion removal will be handled separately

	testAccProvidersAlias = map[string]tf.ResourceProvider{ // want "deprecated tf.ResourceProvider"
		"aws": testAccProviderAlias,
	}

	_ = tf.State{} // keep terraform dependency
}
