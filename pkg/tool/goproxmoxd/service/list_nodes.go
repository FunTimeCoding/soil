package service

import "github.com/luthermonson/go-proxmox"

func (s *Service) ListNodes(instance string) (proxmox.NodeStatuses, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return nil, clientFail
	}

	return c.Nodes()
}
