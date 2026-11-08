package census

import "go/types"

func fieldTypes(n *types.Named) map[string]bool {
	result := map[string]bool{}
	s := n.Underlying().(*types.Struct)

	for i := range s.NumFields() {
		if f := named(s.Field(i).Type()); f != nil {
			result[qualified(f)] = true
		}
	}

	return result
}
