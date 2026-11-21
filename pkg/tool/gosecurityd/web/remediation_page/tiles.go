package remediation_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_result"
	"github.com/funtimecoding/soil/pkg/web/component/tile_grid"
	"maragu.dev/gomponents"
)

func Tiles(r *remediation_result.Result) gomponents.Node {
	return tile_grid.New(
		[]gomponents.Node{
			count(r.Open, "Open items"),
			count(r.Kinds[constant.RemediationUpgrade], "Upgrades to roll out"),
			count(
				r.Kinds[constant.RemediationVulnerability],
				"Vulnerabilities without an upgrade",
			),
			count(r.Kinds[constant.RemediationRelease], "Release upgrades"),
			count(r.Kinds[constant.RemediationReboot], "Reboots"),
			count(
				r.Kinds[constant.RemediationKernel],
				"Kernels whose fixes are out of reach",
			),
			count(r.Addressed, "Addressed by a decision"),
			count(
				r.Unimportant,
				"Vulnerabilities left out: Debian marks them unimportant",
			),
		},
	).Node()
}
