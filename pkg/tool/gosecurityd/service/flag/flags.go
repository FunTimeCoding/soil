package flag

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"

func (s *Service) Flags(includeMitigated bool) ([]flag.Flag, error) {
	return s.store.List(includeMitigated)
}
