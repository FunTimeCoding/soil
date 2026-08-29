package mock_service

import "github.com/luthermonson/go-proxmox"

func (s *Service) GetContainer(
	instance string,
	identifier int,
	node string,
) (*proxmox.Container, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return nil, clientFail
	}

	n, e := c.Node(node)

	if e != nil {
		return nil, e
	}

	return c.Container(n, identifier)
}
