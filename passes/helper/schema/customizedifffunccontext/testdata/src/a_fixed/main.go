package a

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func customizeDiffFunc(ctx context.Context, diff *schema.ResourceDiff, v interface{}) error { // want "add missing context.Context parameter"
	return nil
}

func f() {
	_ = func(ctx context.Context, diff *schema.ResourceDiff, v interface{}) error { // want "add missing context.Context parameter"
		return nil
	}

	_ = context.TODO() // keep context dependency
}
