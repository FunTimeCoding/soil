package revision

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"

func collisions(symbols []*symbol.Symbol) [][]*symbol.Symbol {
	groups := map[string][]*symbol.Symbol{}
	var order []string

	for _, s := range symbols {
		key := s.Key()

		if _, seen := groups[key]; !seen {
			order = append(order, key)
		}

		groups[key] = append(groups[key], s)
	}

	var result [][]*symbol.Symbol

	for _, key := range order {
		if len(groups[key]) > 1 {
			result = append(result, groups[key])
		}
	}

	return result
}
