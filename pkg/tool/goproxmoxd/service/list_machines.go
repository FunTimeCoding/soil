package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListMachines(
	c face.ProxmoxClient,
	node string,
) (proxmox.VirtualMachines, error) {
	nodeNames, e := s.resolveNodeNames(c, node)

	if e != nil {
		return nil, e
	}

	var result proxmox.VirtualMachines

	for _, name := range nodeNames {
		n, e := c.Node(name)

		if e != nil {
			return nil, e
		}

		machines, e := c.Machines(n)

		if e != nil {
			return nil, e
		}

		result = append(result, machines...)
	}

	return result, nil
}
