package report

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"

func byMember(symbols map[string]*symbol.Symbol) map[string][]*symbol.Symbol {
	result := map[string][]*symbol.Symbol{}

	for _, s := range symbols {
		result[s.Member()] = append(result[s.Member()], s)
	}

	return result
}
