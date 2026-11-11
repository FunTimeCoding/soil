package index

import (
	"github.com/funtimecoding/soil/pkg/source/canonical_type"
	"go/types"
)

func Signature(f *types.Func) string {
	return canonical_type.String(
		canonical_type.Unnamed(f.Type().(*types.Signature)),
		qualifier,
	)
}
