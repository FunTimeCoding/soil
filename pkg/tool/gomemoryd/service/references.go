package service

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func (s *Service) References(m *record.Memory) []*reference.Finding {
	if s.referenceRoot == "" || m.ProvenanceFile != "" {
		return nil
	}

	return reference.Check(
		m.Content,
		reference.Bases(m.Metadata[constant.BaseKey]),
		lint.ReferenceResolver(s.referenceRoot),
		s.named,
	)
}
