package service

import "github.com/luthermonson/go-proxmox"

func (s *Service) ListContainerSnapshots(
	instance string,
	identifier int,
	node string,
) ([]*proxmox.ContainerSnapshot, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return nil, clientFail
	}

	ct, e := findContainer(c, identifier, node)

	if e != nil {
		return nil, e
	}

	return c.ContainerSnapshots(ct)
}
