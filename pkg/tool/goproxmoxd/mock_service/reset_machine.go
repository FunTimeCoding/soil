package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) ResetMachine(
	_ face.ProxmoxClient,
	_ int,
	_ string,
) (string, error) {
	return "mock:reset", nil
}
