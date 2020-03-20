package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func f() {
	_ = validation.ValidateJsonString // want "deprecated validation.ValidateJsonString must be replaced with validation.StringIsJSON"
}
