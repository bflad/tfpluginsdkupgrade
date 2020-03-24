package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	tf "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func ProviderAlias() tf.ResourceProvider {
	return &schema.Provider{}
}

func fAlias() {
	_ = func(providers *[]*schema.Provider) map[string]tf.ResourceProviderFactory { // want "tf.ResourceProviderFactory should be replaced"
		return map[string]tf.ResourceProviderFactory{ // want "tf.ResourceProviderFactory should be replaced"
			"aws": func() (tf.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
}
