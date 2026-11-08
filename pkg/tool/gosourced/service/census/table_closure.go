package census

func tableClosure(
	stores map[string]*storeReading,
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

		for q := range s.models {
			if s.read {
				read[tables[q]] = true
			}

			if s.written {
				written[tables[q]] = true
			}
		}

		for _, c := range sortedKeys(s.calls) {
			walk(c)
		}
	}
	walk(key)
}
