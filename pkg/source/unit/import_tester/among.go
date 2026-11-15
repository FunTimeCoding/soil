package import_tester

import "slices"

func Among(names ...string) func(string) bool {
	return func(name string) bool {
		return slices.Contains(names, name)
	}
}
