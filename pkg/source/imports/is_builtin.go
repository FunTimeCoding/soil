package imports

import "go/types"

func isBuiltin(name string) bool {
	return types.Universe.Lookup(name) != nil
}
