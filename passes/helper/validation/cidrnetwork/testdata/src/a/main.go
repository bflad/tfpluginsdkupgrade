package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func f() {
	validation.CIDRNetwork(0, 32) // want "deprecated validation.CIDRNetwork must be replaced with validation.IsCIDRNetwork"
}
