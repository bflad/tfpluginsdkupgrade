package a

import (
	v "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func falias() {
	_ = v.ValidateJsonString // want "deprecated v.ValidateJsonString must be replaced with v.StringIsJSON"
}
