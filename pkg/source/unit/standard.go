package unit

import "slices"

func standard(importPaths ...string) func(string) bool {
	return func(importPath string) bool {
		return slices.Contains(importPaths, importPath)
	}
}
