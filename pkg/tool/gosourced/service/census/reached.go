package census

func reached(
	methods map[string]*reading,
	name string,
) []*reading {
	visited := map[string]bool{name: true}
	result := []*reading{methods[name]}
	var walk func(*reading)
	walk = func(r *reading) {
		for _, c := range sortedKeys(r.calls) {
			if visited[c] || methods[c] == nil {
				continue
			}

			visited[c] = true
			result = append(result, methods[c])
			walk(methods[c])
		}
	}
	walk(methods[name])

	return result
}
