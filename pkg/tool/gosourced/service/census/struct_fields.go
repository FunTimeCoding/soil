package census

import "go/types"

func structFields(n *types.Named) []string {
	s := n.Underlying().(*types.Struct)
	result := make([]string, 0, s.NumFields())

	for i := range s.NumFields() {
		result = append(result, s.Field(i).Name())
	}

	return result
}
