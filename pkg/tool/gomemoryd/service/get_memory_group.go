package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) GetMemoryGroup(
	identifier int64,
) (*store.Memory, []store.Memory, error) {
	parent, e := s.store.GetMemory(identifier)

	if e != nil {
		return nil, nil, e
	}

	summaries, e := s.store.ListChildren(identifier)

	if e != nil {
		return nil, nil, e
	}

	var children []store.Memory

	for _, sum := range summaries {
		m, f := s.store.GetMemory(sum.Identifier)

		if f != nil {
			continue
		}

		children = append(children, *m)
	}

	return parent, children, nil
}
