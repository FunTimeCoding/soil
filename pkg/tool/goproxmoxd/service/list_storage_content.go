package service

import "github.com/luthermonson/go-proxmox"

func (s *Service) ListStorageContent(
	instance string,
	node string,
	storage string,
) ([]*proxmox.StorageContent, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return nil, clientFail
	}

	n, e := c.Node(node)

	if e != nil {
		return nil, e
	}

	s2, e := c.Storage(n, storage)

	if e != nil {
		return nil, e
	}

	return c.StorageContent(s2)
}
