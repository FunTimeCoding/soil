package packager

import (
	"github.com/funtimecoding/soil/pkg/debian"
	"github.com/funtimecoding/soil/pkg/debian/constant"
)

func (p *Packager) SavePostInstall(upgradeMode string) {
	p.saveScript(
		constant.PostInstallScript,
		debian.RenderPostInstall(p.ExecutableName, upgradeMode),
	)
}
