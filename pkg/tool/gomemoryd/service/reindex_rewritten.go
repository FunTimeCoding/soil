package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) reindexRewritten(identifiers []int64) ([]*record.Memory, error) {
	var result []*record.Memory

	for _, identifier := range identifiers {
		m, e := s.store.GetMemory(identifier)

		if e != nil {
			return nil, e
		}

		if e = s.syncIndex(m); e != nil {
			return nil, e
		}

		result = append(result, m)
	}

	return result, nil
}
