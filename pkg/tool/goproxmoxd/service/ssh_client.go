package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) SSHClient(instance string) (face.SnippetClient, error) {
	return s.snippets.Client(instance)
}
