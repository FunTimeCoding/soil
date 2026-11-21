package system

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_result"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_row"
	"github.com/funtimecoding/soil/pkg/vulnerability/identifier"
	"slices"
)

func (s *Service) Affected(name string) (*affected_result.Result, error) {
	if e := identifier.Check(name); e != nil {
		return nil, e
	}

	details, f := s.Details()

	if f != nil {
		return nil, f
	}

	result := affected_result.New(name)

	if c, okay := s.score.Score(name); okay {
		result.Score = c.Score
		result.Severity = c.Severity
	}

	for _, d := range details {
		if v := worst(d.Findings, name); v != nil {
			k := v.Kind()
			result.Count.Add(k)
			result.Rows = append(result.Rows, affected_row.New(d.Summary, v, k))
		}
	}

	slices.SortFunc(
		result.Rows,
		func(a *affected_row.Row, b *affected_row.Row) int {
			return cmp.Or(
				cmp.Compare(
					slices.Index(constant.KindOrder, a.Kind),
					slices.Index(constant.KindOrder, b.Kind),
				),
				cmp.Compare(a.Priority, b.Priority),
				cmp.Compare(a.System, b.System),
			)
		},
	)

	return result, nil
}
