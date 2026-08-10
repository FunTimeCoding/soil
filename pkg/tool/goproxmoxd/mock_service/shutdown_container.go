package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) ShutdownContainer(
	_ face.ProxmoxClient,
	_ int,
	_ string,
) (string, error) {
	return "mock:ct-shutdown", nil
}
