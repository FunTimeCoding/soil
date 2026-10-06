package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) ListDocumentSourced(
	scope string,
) ([]record.SourcedMemory, error) {
	return s.store.ListDocumentSourced(scope)
}
