package host_mapping

func (x *Index) resolveVariable(
	prefix string,
	name string,
) []string {
	c, okay := x.calls[prefix]

	if !okay {
		return nil
	}

	var result []string

	for _, r := range references(c.Expressions[name]) {
		result = append(result, x.resolve(x.parents[prefix], r)...)
	}

	return result
}
