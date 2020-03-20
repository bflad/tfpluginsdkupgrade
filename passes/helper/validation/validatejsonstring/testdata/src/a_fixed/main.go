package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func f() {
	_ = validation.StringIsJSON // want "deprecated validation.ValidateJsonString must be replaced with validation.StringIsJSON"
}
