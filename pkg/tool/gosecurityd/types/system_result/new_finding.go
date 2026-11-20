package system_result

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/status"
)

func NewFinding(
	s *status.Status,
	p *installed.Installed,
) *Finding {
	return &Finding{
		Vulnerability: s.Identifier,
		Package:       p.Name,
		Source:        p.Source,
		Installed:     p.SourceVersion,
		Candidate:     p.Candidate,
		Urgency:       s.Urgency,
		Fix:           constant.FindingNone,
	}
}
