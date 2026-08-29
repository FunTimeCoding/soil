package mock_service

import "github.com/luthermonson/go-proxmox"

func (s *Service) ListContainerSnapshots(
	_ string,
	_ int,
	_ string,
) ([]*proxmox.ContainerSnapshot, error) {
	return nil, nil
}
