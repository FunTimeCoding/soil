package census

func throughFunctions(
	r *reading,
	functions map[string]*reading,
) {
	visited := map[string]bool{}
	var walk func(*reading)
	walk = func(x *reading) {
		for _, g := range sortedKeys(x.functions) {
			if visited[g] || functions[g] == nil {
				continue
			}

			visited[g] = true
			merge(r, functions[g])
			walk(functions[g])
		}
	}
	walk(r)

	for g := range visited {
		r.functions[g] = true
	}
}
