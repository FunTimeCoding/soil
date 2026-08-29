package mock_service

import "github.com/luthermonson/go-proxmox"

func (s *Service) ListContainers(
	instance string,
	node string,
) (proxmox.Containers, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return nil, clientFail
	}

	if node == "" {
		return nil, nil
	}

	n, e := c.Node(node)

	if e != nil {
		return nil, e
	}

	return c.Containers(n)
}
