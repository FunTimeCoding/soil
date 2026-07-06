package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/resolve_result"

func (s *Service) ResolveSessionIdentifier(query string) (*resolve_result.Result, error) {
	return s.store.ResolveSessionIdentifier(query)
}
