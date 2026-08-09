package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (s *Service) UpdateMemory(
	identifier int64,
	o *save_option.Option,
) (*store.Memory, error) {
	existing, e := s.store.GetMemory(identifier)

	if e != nil {
		return nil, e
	}

	if o.Tags == nil {
		o.Tags = existing.Tags
	}

	if o.Metadata == nil {
		o.Metadata = existing.Metadata
	}

	if o.ProvenanceHash == "" {
		o.ProvenanceHash = existing.ProvenanceHash
	}

	if o.Ordinal == 0 {
		o.Ordinal = existing.Ordinal
	}

	if e = s.store.UpdateMemory(identifier, o); e != nil {
		return nil, e
	}

	m, e := s.store.GetMemory(identifier)

	if e != nil {
		return nil, e
	}

	if e = s.syncIndex(m); e != nil {
		return nil, e
	}

	return m, nil
}
