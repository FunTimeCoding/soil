package peered

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/proxmox/face"
)

func New(
	c face.ProxmoxClient,
	g *peer.Guard,
) *Peered {
	return &Peered{client: c, guard: g}
}
