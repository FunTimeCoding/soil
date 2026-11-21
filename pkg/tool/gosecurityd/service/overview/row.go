package overview

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/subject"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
)

func (s *Service) row(f *flag.Flag) *overview_result.Row {
	if subject.Kind(f.Subject) != constant.RemediationVulnerability {
		return overview_result.NewRow(f, nil, nil, "")
	}

	d, e := s.debian.Status(f.Subject, "")

	return overview_result.NewRow(f, d, e, fix(d, e))
}
