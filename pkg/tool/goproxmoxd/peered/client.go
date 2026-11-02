package peered

import "github.com/funtimecoding/soil/pkg/proxmox/face"

func (p *Peered) Client() face.ProxmoxClient {
	return p.client
}
