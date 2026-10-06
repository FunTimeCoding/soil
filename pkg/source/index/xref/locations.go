package xref

import (
	"fmt"
	"path/filepath"
	"sort"
)

func (i *Index) Locations(target string) []string {
	var result []string

	for unit, r := range i.units {
		for _, s := range r.Targets[target] {
			result = append(
				result,
				fmt.Sprintf(
					"%s:%d:%d",
					filepath.Join(i.directories[unit], s.File),
					s.Line,
					s.Column,
				),
			)
		}
	}

	sort.Strings(result)

	return result
}
