package service

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

	return findContainer(c, identifier, node)
}
