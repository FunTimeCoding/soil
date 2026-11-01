package entry

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"

func New(
	class string,
	from *symbol.Symbol,
	to *symbol.Symbol,
) *Entry {
	return &Entry{Class: class, From: from, To: to}
}
