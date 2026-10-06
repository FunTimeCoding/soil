package service

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"
)

func (s *Service) ReferenceCensus() ([]*reference.Report, error) {
	if s.referenceRoot == "" {
		return nil, nil
	}

	memories, e := s.ListMemoriesWithContent("", constant.AllScope)

	if e != nil {
		return nil, e
	}

	r := lint.ReferenceResolver(s.referenceRoot)
	var result []*reference.Report

	for _, m := range memories {
		if m.ProvenanceFile != "" {
			continue
		}

		f := reference.Check(
			m.Content,
			reference.Bases(m.Metadata[constant.BaseKey]),
			r,
			s.named,
		)

		if len(f) > 0 {
			result = append(
				result,
				reference.NewReport(m.Identifier, m.Name, f),
			)
		}
	}

	return result, nil
}
