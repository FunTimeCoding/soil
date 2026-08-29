package service

import "github.com/luthermonson/go-proxmox"

func (s *Service) ListContainers(
	instance string,
	node string,
) (proxmox.Containers, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return nil, clientFail
	}

	nodeNames, e := s.resolveNodeNames(c, node)

	if e != nil {
		return nil, e
	}

	var result proxmox.Containers

	for _, name := range nodeNames {
		n, e := c.Node(name)

		if e != nil {
			return nil, e
		}

		containers, e := c.Containers(n)

		if e != nil {
			return nil, e
		}

		result = append(result, containers...)
	}

	return result, nil
}
