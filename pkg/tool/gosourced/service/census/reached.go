package census

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/census/reading"

func reached(
	methods map[string]*reading.Reading,
	name string,
) []*reading.Reading {
	visited := map[string]bool{name: true}
	result := []*reading.Reading{methods[name]}
	var walk func(*reading.Reading)
	walk = func(r *reading.Reading) {
		for _, c := range sortedKeys(r.Calls) {
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
