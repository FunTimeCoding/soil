package mock_service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListContainerSnapshots(
	_ face.ProxmoxClient,
	_ int,
	_ string,
) ([]*proxmox.ContainerSnapshot, error) {
	return nil, nil
}
