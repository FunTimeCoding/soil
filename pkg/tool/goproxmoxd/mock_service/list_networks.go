package mock_service

import "github.com/luthermonson/go-proxmox"

func (s *Service) ListNetworks(
	_ string,
	_ string,
) (proxmox.NodeNetworks, error) {
	return nil, nil
}
