package report

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
)

func leftovers(
	before map[string]*symbol.Symbol,
	after map[string]*symbol.Symbol,
) []*entry.Entry {
	var result []*entry.Entry

	for _, s := range before {
		result = append(result, entry.New(constant.ClassRemoved, s, nil))
	}

	for _, s := range after {
		result = append(result, entry.New(constant.ClassAdded, nil, s))
	}

	return result
}
