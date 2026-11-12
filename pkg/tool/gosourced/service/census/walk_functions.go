package census

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/census/reading"

func walkFunctions(
	r *reading.Reading,
	functions map[string]*reading.Reading,
) {
	visited := map[string]bool{}
	var walk func(*reading.Reading)
	walk = func(x *reading.Reading) {
		for _, g := range sortedKeys(x.Functions) {
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
		r.Functions[g] = true
	}
}
