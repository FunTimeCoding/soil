package symbol

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/token"
)

func movedContains(
	entries []*relocation.Entry,
	position token.Pos,
) bool {
	for _, entry := range entries {
		if position >= entry.Node.Pos() && position <= entry.Node.End() {
			return true
		}
	}

	return false
}
