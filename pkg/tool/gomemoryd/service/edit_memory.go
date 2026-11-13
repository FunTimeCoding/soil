package service

import (
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (s *Service) EditMemory(
	identifier int64,
	o *save_option.Option,
) (*record.Memory, []*record.Memory, error) {
	existing, e := s.store.GetMemory(identifier)

	if e != nil {
		return nil, nil, e
	}

	if existing.ProvenanceFile != "" {
		return nil, nil, forbidden.Format(
			"memory %d is document-sourced and read-only; edit %s",
			identifier,
			existing.ProvenanceFile,
		)
	}

	return s.UpdateMemory(identifier, o)
}
