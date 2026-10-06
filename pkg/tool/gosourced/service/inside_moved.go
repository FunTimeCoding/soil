package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"go/token"
)

func insideMoved(
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
