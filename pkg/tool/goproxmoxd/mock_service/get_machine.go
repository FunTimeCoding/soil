package mock_service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) GetMachine(
	c face.ProxmoxClient,
	identifier int,
	node string,
) (*proxmox.VirtualMachine, error) {
	n, e := c.Node(node)

	if e != nil {
		return nil, e
	}

	return c.Machine(n, identifier)
}
