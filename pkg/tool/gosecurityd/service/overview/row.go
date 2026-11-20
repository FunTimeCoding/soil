package overview

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
)

func (s *Service) row(f *flag.Flag) *overview_result.Row {
	d, e := s.debian.Status(f.Vulnerability, "")

	return overview_result.NewRow(f, d, e, fix(d, e))
}
