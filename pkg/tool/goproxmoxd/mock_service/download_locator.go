package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) DownloadLocator(
	_ face.ProxmoxClient,
	_ string,
	_ string,
	_ string,
	_ string,
	_ string,
) error {
	return nil
}
