package pointer_tester

import "slices"

func exists(all []string) func(string) bool {
	return func(p string) bool {
		return slices.Contains(all, p)
	}
}
