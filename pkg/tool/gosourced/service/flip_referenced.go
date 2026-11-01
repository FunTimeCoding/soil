package service

import (
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"unicode"
)

func flipReferenced(
	entries []*relocation.Entry,
	references map[*relocation.Entry][]resolve_reference.Reference,
) {
	for _, entry := range entries {
		if !unicode.IsLower(rune(entry.Symbol[0])) {
			continue
		}

		for _, f := range references[entry] {
			if insideMoved(entries, f.Ident.Pos()) {
				continue
			}

			entry.Flipped = true
			entry.NewName = FlipName(entry.Symbol)

			break
		}
	}
}
