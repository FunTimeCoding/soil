package service

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func (s *Service) References(m *store.Memory) []*reference.Finding {
	if s.referenceRoot == "" || m.ProvenanceFile != "" {
		return nil
	}

	return reference.Check(
		m.Content,
		lint.ReferenceResolver(s.referenceRoot),
		s.named,
	)
}
