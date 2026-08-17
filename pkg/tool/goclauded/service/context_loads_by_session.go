package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"

func (s *Service) ContextLoadsBySession(
	sessionIdentifier string,
) ([]context_load.Load, error) {
	return s.store.ContextLoadsBySession(sessionIdentifier)
}
