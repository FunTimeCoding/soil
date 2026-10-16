package host_mapping

func (x *Index) hosts(
	address string,
	visited map[string]bool,
) []string {
	if visited[address] {
		return nil
	}

	visited[address] = true
	r, okay := x.resources[address]

	if !okay {
		return nil
	}

	if attribute, named := x.attribute[r.Type]; named {
		var result []string

		for _, v := range x.values[address] {
			if s, isString := v[attribute].(string); isString && s != "" {
				result = append(result, s)
			}
		}

		return result
	}

	var result []string

	for _, reference := range references(r.Expressions) {
		for _, target := range x.resolve(x.scope[address], reference) {
			result = append(result, x.hosts(target, visited)...)
		}
	}

	return result
}
