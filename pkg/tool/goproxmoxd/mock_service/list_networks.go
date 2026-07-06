package mock_service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListNetworks(
	_ face.ProxmoxClient,
	_ string,
) (proxmox.NodeNetworks, error) {
	return nil, nil
}
