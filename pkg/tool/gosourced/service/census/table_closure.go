package census

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/census/reading"

func tableClosure(
	stores map[string]*reading.Store,
	tables map[string]string,
	key string,
	read map[string]bool,
	written map[string]bool,
) {
	visited := map[string]bool{}
	var walk func(string)
	walk = func(k string) {
		s := stores[k]

		if visited[k] || s == nil {
			return
		}

		visited[k] = true

		for q := range s.Models {
			if s.Read {
				read[tables[q]] = true
			}

			if s.Written {
				written[tables[q]] = true
			}
		}

		for _, c := range sortedKeys(s.Calls) {
			walk(c)
		}
	}
	walk(key)
}
