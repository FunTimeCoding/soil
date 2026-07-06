package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) SSHClient(_ string) (face.SnippetClient, error) {
	return s.snippetClient, nil
}
