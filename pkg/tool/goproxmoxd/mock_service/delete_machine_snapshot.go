package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) DeleteMachineSnapshot(
	_ face.ProxmoxClient,
	_ int,
	_ string,
	_ string,
) (string, error) {
	return "mock:delete-snapshot", nil
}
