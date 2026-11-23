package remediation

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/minimum"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/subject"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_builder"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_result"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"slices"
)

func (s *Service) List(
	selection *system_selection.Selection,
) (*remediation_result.Result, error) {
	m := selection.Minimum

	if e := minimum.Check(m); e != nil {
		return nil, e
	}

	details, f := s.system.Selected(selection.WithMinimum(0))

	if f != nil {
		return nil, f
	}

	b := remediation_builder.New()

	for _, d := range details {
		if d.Summary.Coverage == constant.CoverageUntracked {
			b.Item(
				subject.Host(constant.RemediationRelease, d.Summary.Name),
				d.Summary,
				nil,
			).Unknown = true

			continue
		}

		for _, v := range d.Findings {
			if v.Score < m {
				continue
			}

			if key := classify(d.Summary, v); key != "" {
				b.Add(key, d.Summary, v)
			} else {
				b.LeaveOut(v.Vulnerability)
			}
		}
	}

	flags, g := s.flag.Flags(true)

	if g != nil {
		return nil, g
	}

	result := b.Result(flags)
	slices.SortFunc(
		result.Items,
		func(a *remediation_item.Item, b *remediation_item.Item) int {
			return cmp.Or(
				cmp.Compare(rank(a), rank(b)),
				cmp.Compare(
					len(b.Hosts)*len(b.Vulnerabilities),
					len(a.Hosts)*len(a.Vulnerabilities),
				),
				cmp.Compare(len(b.Hosts), len(a.Hosts)),
				cmp.Compare(a.Subject, b.Subject),
			)
		},
	)

	return result, nil
}
