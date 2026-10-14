package unclosed_resource

import "go/types"

func fieldName(
	t types.Type,
	index int,
) string {
	s, okay := dereference(t).Underlying().(*types.Struct)

	if !okay || index >= s.NumFields() {
		return ""
	}

	return s.Field(index).Name()
}
