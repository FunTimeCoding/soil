package module

import (
	"github.com/funtimecoding/soil/pkg/source/canonical_type"
	"go/types"
)

func nameless(t types.Type) types.Type {
	if s, okay := t.(*types.Signature); okay {
		return canonical_type.Unnamed(s)
	}

	return t
}
