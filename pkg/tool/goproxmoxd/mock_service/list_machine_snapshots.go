package mock_service

import "github.com/luthermonson/go-proxmox"

func (s *Service) ListMachineSnapshots(
	_ string,
	_ int,
	_ string,
) ([]*proxmox.VirtualMachineSnapshot, error) {
	return nil, nil
}
