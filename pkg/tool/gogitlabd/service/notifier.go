package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) Notifier() face.EventNotifier {
	return s.notifier
}
