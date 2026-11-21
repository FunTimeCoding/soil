package system_result

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/status"
)

func NewKernelFinding(
	s *status.Status,
	running *installed.Installed,
	backports bool,
) *Finding {
	return &Finding{
		Vulnerability: s.Identifier,
		Package:       running.Name,
		Source:        constant.KernelSource,
		Installed:     running.Version,
		Urgency:       s.Urgency,
		Fix:           constant.FindingNone,
		Kernel:        true,
		Backports:     backports,
	}
}
