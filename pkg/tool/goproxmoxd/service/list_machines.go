package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListMachines(
	instance string,
	node string,
) (proxmox.VirtualMachines, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (proxmox.VirtualMachines, error) {
			nodeNames, e := s.resolveNodeNames(c, node)

			if e != nil {
				return nil, e
			}

			var result proxmox.VirtualMachines

			for _, name := range nodeNames {
				n, f := c.Node(name)

				if f != nil {
					return nil, f
				}

				machines, f := c.Machines(n)

				if f != nil {
					return nil, f
				}

				result = append(result, machines...)
			}

			return result, nil
		},
	)
}
