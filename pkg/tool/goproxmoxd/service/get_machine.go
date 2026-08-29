package service

import "github.com/luthermonson/go-proxmox"

func (s *Service) GetMachine(
	instance string,
	identifier int,
	node string,
) (*proxmox.VirtualMachine, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return nil, clientFail
	}

	return findMachine(c, identifier, node)
}
