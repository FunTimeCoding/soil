package peered

import "github.com/funtimecoding/soil/pkg/peer"

func (p *Peered) Guard() *peer.Guard {
	return p.guard
}
