package relocator

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"go/token"
)

func MovedContains(
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
