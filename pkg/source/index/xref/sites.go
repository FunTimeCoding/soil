package xref

func (i *Index) Sites(target string) map[string][]*Site {
	result := make(map[string][]*Site)

	for unit, r := range i.units {
		if sites := r.Targets[target]; len(sites) > 0 {
			result[unit] = sites
		}
	}

	return result
}
