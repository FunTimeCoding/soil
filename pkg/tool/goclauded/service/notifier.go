package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/face"

func (s *Service) Notifier() face.Notifier {
	return s.notifier
}
