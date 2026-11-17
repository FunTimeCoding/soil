package service_tester

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"

func WithMemoryPeer(c face.Client) Option {
	return func(s *setup) {
		s.memoryPeer = c
	}
}
