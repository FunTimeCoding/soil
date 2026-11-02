package peered

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/proxmox/face"
)

type Peered struct {
	client face.ProxmoxClient
	guard  *peer.Guard
}
