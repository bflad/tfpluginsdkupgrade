package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func ProviderAlias() terraform.ResourceProvider {
	return &s.Provider{}
}

func fAlias() {
	_ = ProviderAlias().(*s.Provider) // want "\\.\\(\\*s.Provider\\) type assertion should be removed"

	_ = ProviderAlias().(*s.Provider).InternalValidate() // want "\\.\\(\\*s.Provider\\) type assertion should be removed"

	_ = func(providers *[]*s.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"aws": func() (terraform.ResourceProvider, error) {
				p := ProviderAlias()
				*providers = append(*providers, p.(*s.Provider)) // want "\\.\\(\\*s.Provider\\) type assertion should be removed"
				return p, nil
			},
		}
	}
}
