package relocator

import (
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"unicode"
)

func FlipReferenced(
	entries []*relocation.Entry,
	references map[*relocation.Entry][]resolve_reference.Reference,
) {
	for _, entry := range entries {
		if !unicode.IsLower(rune(entry.Symbol[0])) {
			continue
		}

		for _, f := range references[entry] {
			if MovedContains(entries, f.Ident.Pos()) {
				continue
			}

			entry.Flipped = true
			entry.NewName = identifier.FlipName(entry.Symbol)

			break
		}
	}
}
