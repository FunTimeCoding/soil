package peered

import "github.com/funtimecoding/soil/pkg/peer"

func (p *Peered) AlertmanagerGuard() *peer.Guard {
	return p.alertmanagerGuard
}
