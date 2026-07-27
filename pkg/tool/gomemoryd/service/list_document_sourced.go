package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) ListDocumentSourced(
	scope string,
) ([]store.SourcedMemory, error) {
	return s.store.ListDocumentSourced(scope)
}
