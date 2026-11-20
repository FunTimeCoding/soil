package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/tracker"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/version"
)

func match(
	t *tracker.Tracker,
	release string,
	p *installed.Installed,
	supported bool,
) []*system_result.Finding {
	var result []*system_result.Finding

	for _, s := range t.Package(p.Source, release) {
		f := system_result.NewFinding(s, p)

		if s.Status == vulnerabilityConstant.DebianResolved {
			if s.Fixed == "" ||
				version.Compare(p.SourceVersion, s.Fixed) >= 0 {
				continue
			}

			f.Fixed = s.Fixed
			f.Upgradable = p.Candidate != "" &&
				version.Compare(p.Candidate, s.Fixed) >= 0
			f.Fix = constant.FindingPaid

			if supported || f.Upgradable {
				f.Fix = constant.FindingDebian
			}
		}

		result = append(result, f)
	}

	return result
}
