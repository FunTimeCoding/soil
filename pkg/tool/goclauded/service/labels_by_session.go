package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"

func (s *Service) LabelsBySession(sessionIdentifier string) ([]label.Label, error) {
	return s.store.LabelsBySession(sessionIdentifier)
}
