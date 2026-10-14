package unclosed_resource

import "go/types"

func dereference(t types.Type) types.Type {
	if p, okay := t.(*types.Pointer); okay {
		return p.Elem()
	}

	return t
}
