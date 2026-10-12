package unit

import "slices"

func taken(names ...string) func(string) bool {
	return func(name string) bool {
		return slices.Contains(names, name)
	}
}
