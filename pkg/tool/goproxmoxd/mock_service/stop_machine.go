package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) StopMachine(
	_ face.ProxmoxClient,
	_ int,
	_ string,
) (string, error) {
	return "mock:stop", nil
}
