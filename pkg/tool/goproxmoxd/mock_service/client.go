package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) Client(_ string) (face.ProxmoxClient, error) {
	return s.proxmoxClient, nil
}
