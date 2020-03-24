package a

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func customizeDiffFunc(diff *schema.ResourceDiff, v interface{}) error { // want "add missing context.Context parameter"
	return nil
}

func f() {
	_ = func(diff *schema.ResourceDiff, v interface{}) error { // want "add missing context.Context parameter"
		return nil
	}

	_ = context.TODO() // keep context dependency
}
