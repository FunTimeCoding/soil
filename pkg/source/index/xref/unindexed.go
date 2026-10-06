package xref

import "sort"

func (i *Index) Unindexed() []string {
	var result []string

	for unit := range i.directories {
		if _, okay := i.units[unit]; !okay {
			result = append(result, unit)
		}
	}

	sort.Strings(result)

	return result
}
