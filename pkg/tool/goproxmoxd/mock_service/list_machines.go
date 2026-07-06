package mock_service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListMachines(
	c face.ProxmoxClient,
	node string,
) (proxmox.VirtualMachines, error) {
	if node == "" {
		return nil, nil
	}

	n, e := c.Node(node)

	if e != nil {
		return nil, e
	}

	return c.Machines(n)
}
