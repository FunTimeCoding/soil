package generation

import (
	"os"
	"slices"
)

func siblings(
	root string,
	name string,
) []string {
	entries, e := os.ReadDir(root)

	if e != nil {
		return nil
	}

	var result []string

	for _, v := range entries {
		if v.Name() != name {
			result = append(result, v.Name())
		}
	}

	slices.Sort(result)

	return result
}
