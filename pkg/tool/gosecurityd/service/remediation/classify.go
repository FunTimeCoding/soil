package remediation

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/subject"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
)

func classify(
	s *system_result.Summary,
	f *system_result.Finding,
) string {
	switch {
	case f.Urgency == constant.Unimportant:
		return ""
	case s.Coverage == constant.CoverageTracked:
		return subject.Host(constant.RemediationRelease, s.Name)
	case f.Reboot:
		return subject.Host(constant.RemediationReboot, s.Name)
	case f.Upgradable:
		return subject.Upgrade(s.Release, f.Label())
	case f.UnreachableKernel():
		return subject.Host(constant.RemediationKernel, s.Name)
	}

	return f.Vulnerability
}
