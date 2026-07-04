package service

import "github.com/funtimecoding/go-library/pkg/face"

func (s *Service) Notifier() face.EventNotifier {
	return s.notifier
}
