package report

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"

func byKey(symbols []*symbol.Symbol) map[string]*symbol.Symbol {
	result := map[string]*symbol.Symbol{}

	for _, s := range symbols {
		result[s.Key()] = s
	}

	return result
}
