package service

import "github.com/luthermonson/go-proxmox"

func (s *Service) ListMachines(
	instance string,
	node string,
) (proxmox.VirtualMachines, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return nil, clientFail
	}

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
