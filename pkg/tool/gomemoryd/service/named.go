package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func (s *Service) named(
	scope string,
	name string,
) bool {
	if scope == constant.DefaultScope {
		scope = ""
	}

	found, e := s.store.MemoryNamed(scope, name)

	return e == nil && found
}
