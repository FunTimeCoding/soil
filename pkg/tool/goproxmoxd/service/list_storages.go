package service

import "github.com/luthermonson/go-proxmox"

func (s *Service) ListStorages(
	instance string,
	node string,
) (proxmox.Storages, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return nil, clientFail
	}

	n, e := c.Node(node)

	if e != nil {
		return nil, e
	}

	return c.Storages(n)
}
