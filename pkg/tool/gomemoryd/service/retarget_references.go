package service

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/pointer/retarget"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/retarget_report"
)

func (s *Service) RetargetReferences(
	moves *move_set.Set,
	dryRun bool,
) ([]*retarget_report.Report, error) {
	if s.referenceRoot == "" {
		return nil, nil
	}

	memories, e := s.ListMemoriesWithContent("", constant.AllScope)

	if e != nil {
		return nil, e
	}

	set := relativeSet(s.referenceRoot, moves)
	r := retarget.Through(lint.ReferenceResolver(s.referenceRoot), set.Paths)
	var result []*retarget_report.Report

	for i := range memories {
		m := &memories[i]
		base := m.Metadata[constant.BaseKey]

		if m.ProvenanceFile != "" ||
			!retarget.Mentions(base, m.Content, set) {
			continue
		}

		content, movedBase, changes := retargetMemory(r, set, m)

		if len(changes) == 0 {
			continue
		}

		result = append(
			result,
			retarget_report.New(m.Identifier, m.Name, changes),
		)

		if dryRun || (content == m.Content && movedBase == base) {
			continue
		}

		o := save_option.New()
		o.Name = m.Name
		o.Content = content
		o.Description = m.Description

		if movedBase != base {
			o.Base = &movedBase
		}

		if _, _, f := s.UpdateMemory(m.Identifier, o); f != nil {
			return nil, f
		}
	}

	return result, nil
}
