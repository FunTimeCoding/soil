package packager

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/debian"
)

func (p *Packager) SavePostInstall(upgradeMode string) {
	p.saveScript(
		constant.DebianPostInstallScript,
		debian.RenderPostInstall(p.ExecutableName, upgradeMode),
	)
}
