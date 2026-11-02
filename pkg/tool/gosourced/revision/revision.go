package revision

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"

type Revision struct {
	Module   string
	Symbols  []*symbol.Symbol
	Unparsed []string
}
