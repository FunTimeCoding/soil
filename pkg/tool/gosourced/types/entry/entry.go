package entry

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"

type Entry struct {
	Class string
	From  *symbol.Symbol
	To    *symbol.Symbol
}
