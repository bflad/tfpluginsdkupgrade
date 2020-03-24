package a

import (
	"context"

	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func customizeDiffFuncAlias(ctx context.Context, diff *s.ResourceDiff, v interface{}) error { // want "add missing context.Context parameter"
	return nil
}

func fAlias() {
	_ = func(ctx context.Context, diff *s.ResourceDiff, v interface{}) error { // want "add missing context.Context parameter"
		return nil
	}

	_ = context.TODO() // keep context dependency
}
