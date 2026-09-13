package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"

func (s *Service) LabelsBySessions(
	sessionIdentifiers []string,
) (map[string][]label.Label, error) {
	return s.store.LabelsBySessions(sessionIdentifiers)
}
