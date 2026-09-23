package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/store/frame"

func (s *Service) Frames(session string) []*frame.Frame {
	return s.store.Frames(session)
}
