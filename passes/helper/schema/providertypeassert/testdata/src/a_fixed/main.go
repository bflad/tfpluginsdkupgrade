package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{}
}

func f() {
	_ = Provider() // want "\\.\\(\\*schema.Provider\\) type assertion should be removed"

	_ = Provider().InternalValidate() // want "\\.\\(\\*schema.Provider\\) type assertion should be removed"

	_ = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"aws": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p) // want "\\.\\(\\*schema.Provider\\) type assertion should be removed"
				return p, nil
			},
		}
	}
}
