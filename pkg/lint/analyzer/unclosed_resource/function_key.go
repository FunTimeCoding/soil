package unclosed_resource

import "go/types"

func functionKey(f *types.Func) string {
	return f.Origin().FullName()
}
