package system

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"

func (s *Service) Detail(name string) (*system_result.Detail, error) {
	v, e := s.store.SystemByName(name)

	if e != nil {
		return nil, e
	}

	return s.assess(v)
}
