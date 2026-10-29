package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) kill(v face.Stream) {
	if e := v.Kill(); e != nil {
		s.logger.Structured("kill log stream failed", "error", e.Error())
	}
}
