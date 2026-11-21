package system

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/version"
	"strings"
)

func (s *Service) matchKernel(
	v *system.System,
	packages []installed.Installed,
) ([]*system_result.Finding, string, string, bool) {
	var running *installed.Installed
	newest := ""
	candidate := ""

	for i := range packages {
		p := &packages[i]

		if !strings.HasPrefix(p.Name, constant.KernelImagePrefix) {
			continue
		}

		if v.Kernel != "" &&
			p.Name == join.Empty(constant.KernelImagePrefix, v.Kernel) {
			running = p
		}

		base, _ := backports(p.Version)
		newest = highest(newest, base)

		if p.Candidate != "" {
			c, _ := backports(p.Candidate)
			candidate = highest(candidate, c)
		}
	}

	if running == nil {
		return nil, "", newest, false
	}

	base, number := backports(running.Version)
	release := v.Release

	if number > 0 {
		if after := releaseAfter(number); after != "" {
			release = after
		}
	}

	source := s.debian.Covering(release)

	if source == nil || source.Tracker() == nil {
		return nil, running.Version, newest, false
	}

	var result []*system_result.Finding

	for _, st := range source.Tracker().Package(constant.KernelSource, release) {
		f := system_result.NewKernelFinding(st, running, number > 0)

		if st.Status == vulnerabilityConstant.DebianResolved {
			if st.Fixed == "" || version.Compare(base, st.Fixed) >= 0 {
				continue
			}

			f.Fixed = st.Fixed
			f.Candidate = candidate
			f.Reboot = version.Compare(newest, st.Fixed) >= 0
			f.Upgradable = candidate != "" &&
				version.Compare(candidate, st.Fixed) >= 0
			f.Fix = constant.FindingPaid

			if source.Supported() || f.Reboot || f.Upgradable {
				f.Fix = constant.FindingDebian
			}
		}

		result = append(result, f)
	}

	return result, running.Version, newest, true
}
