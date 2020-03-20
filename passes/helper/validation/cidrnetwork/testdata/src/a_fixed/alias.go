package a

import (
	v "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func falias() {
	v.IsCIDRNetwork(0, 32) // want "deprecated v.CIDRNetwork must be replaced with v.IsCIDRNetwork"
}
