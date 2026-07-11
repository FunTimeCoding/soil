package service

import "go/token"

func insideMoved(
	entries []*moveEntry,
	position token.Pos,
) bool {
	for _, entry := range entries {
		if position >= entry.node.Pos() && position <= entry.node.End() {
			return true
		}
	}

	return false
}
