package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/report"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/revision"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/module_range"
)

func (s *Service) crossReport(
	before map[string][]byte,
	after map[string][]byte,
	ranges []*module_range.Range,
) (*report.Report, error) {
	from := []*revision.Revision{revision.New(before)}
	to := []*revision.Revision{revision.New(after)}

	for _, r := range ranges {
		m, _ := s.Module(r.Name)
		was, e := revisionFiles(m.Directory, r.From)

		if e != nil {
			return nil, e
		}

		now, f := revisionFiles(m.Directory, r.To)

		if f != nil {
			return nil, f
		}

		from = append(from, revision.New(was))
		to = append(to, revision.New(now))
	}

	own := []string{from[0].Module, to[0].Module}
	result := report.New(revision.Merge(from...), revision.Merge(to...))
	result.Modules = ranges
	result.Keep(own...)

	return result, nil
}
