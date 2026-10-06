package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (s *Service) UpdateMemory(
	identifier int64,
	o *save_option.Option,
) (*record.Memory, []*record.Memory, error) {
	existing, e := s.store.GetMemory(identifier)

	if e != nil {
		return nil, nil, e
	}

	keepStored(o, existing)
	rewritten, e := s.store.UpdateMemory(identifier, o)

	if e != nil {
		return nil, nil, e
	}

	m, e := s.store.GetMemory(identifier)

	if e != nil {
		return nil, nil, e
	}

	if e = s.syncIndex(m); e != nil {
		return nil, nil, e
	}

	citing, e := s.reindexRewritten(rewritten)

	if e != nil {
		return nil, nil, e
	}

	return m, citing, nil
}
