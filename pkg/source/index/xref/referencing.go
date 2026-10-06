package xref

import "sort"

func (i *Index) Referencing(target string) []string {
	var result []string

	for unit, r := range i.units {
		if len(r.Targets[target]) > 0 {
			result = append(result, unit)
		}
	}

	sort.Strings(result)

	return result
}
