package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) SetClient(
	instance string,
	c face.ProxmoxClient,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.clients[instance] = c
}
