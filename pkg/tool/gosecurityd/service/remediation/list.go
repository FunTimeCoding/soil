package remediation

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/minimum"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/subject"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_builder"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_result"
	"slices"
)

func (s *Service) List(m float64) (*remediation_result.Result, error) {
	if e := minimum.Check(m); e != nil {
		return nil, e
	}

	details, e := s.system.Details()

	if e != nil {
		return nil, e
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

	flags, f := s.flag.Flags(true)

	if f != nil {
		return nil, f
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
