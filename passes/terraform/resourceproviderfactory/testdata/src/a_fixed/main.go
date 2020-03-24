package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{}
}

func f() {
	_ = func(providers *[]*schema.Provider) map[string]func() (*schema.Provider, error) { // want "terraform.ResourceProviderFactory should be replaced"
		return map[string]func() (*schema.Provider, error){ // want "terraform.ResourceProviderFactory should be replaced"
			"aws": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
}
