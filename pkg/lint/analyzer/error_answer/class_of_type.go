package error_answer

import "go/types"

func classOfType(t types.Type) string {
	for {
		pointer, okay := t.(*types.Pointer)

		if !okay {
			break
		}

		t = pointer.Elem()
	}

	named, okay := t.(*types.Named)

	if !okay || named.Obj().Pkg() == nil {
		return ""
	}

	return knownClass(named.Obj().Pkg().Path())
}
