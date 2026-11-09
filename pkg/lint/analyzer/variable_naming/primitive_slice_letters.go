package variable_naming

import "go/types"

func primitiveSliceLetters(t types.Type) []string {
	s, okay := t.Underlying().(*types.Slice)

	if !okay {
		return nil
	}

	return typeLetters(s.Elem())
}
